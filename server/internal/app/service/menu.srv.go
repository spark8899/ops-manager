package service

import (
	"context"
	"fmt"
	"os"

	"github.com/google/wire"

	"github.com/spark8899/ops-manager/server/internal/app/contextx"
	"github.com/spark8899/ops-manager/server/internal/app/dao"
	"github.com/spark8899/ops-manager/server/internal/app/schema"
	"github.com/spark8899/ops-manager/server/pkg/errors"
	"github.com/spark8899/ops-manager/server/pkg/util/snowflake"
	"github.com/spark8899/ops-manager/server/pkg/util/yaml"
)

var MenuSet = wire.NewSet(wire.Struct(new(MenuSrv), "*"))

type MenuSrv struct {
	TransRepo              *dao.TransRepo
	MenuRepo               *dao.MenuRepo
	MenuActionRepo         *dao.MenuActionRepo
	MenuActionResourceRepo *dao.MenuActionResourceRepo
}

func (a *MenuSrv) InitData(ctx context.Context, dataFile string) error {
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return nil
	}

	data, err := a.readData(dataFile)
	if err != nil {
		return err
	}

	return a.createMenus(ctx, 0, data)
}

func (a *MenuSrv) readData(name string) (schema.MenuTrees, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data schema.MenuTrees
	d := yaml.NewDecoder(file)
	d.SetStrict(true)
	err = d.Decode(&data)
	return data, err
}

func (a *MenuSrv) createMenus(ctx context.Context, parentID uint64, list schema.MenuTrees) error {
	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		for _, item := range list {
			sitem := schema.Menu{
				Name:     item.Name,
				Sequence: item.Sequence,
				Icon:     item.Icon,
				Router:   item.Router,
				ParentID: parentID,
				Status:   1,
				IsShow:   1,
				Actions:  item.Actions,
			}
			if v := item.IsShow; v > 0 {
				sitem.IsShow = v
			}

			nsitem, err := a.Create(ctx, sitem)
			if err != nil {
				return err
			}

			if item.Children != nil && len(*item.Children) > 0 {
				err := a.createMenus(ctx, nsitem.ID, *item.Children)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})
}

func (a *MenuSrv) Query(ctx context.Context, params schema.MenuQueryParam, opts ...schema.MenuQueryOptions) (*schema.MenuQueryResult, error) {
	menuActionResult, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{})
	if err != nil {
		return nil, err
	}

	result, err := a.MenuRepo.Query(ctx, params, opts...)
	if err != nil {
		return nil, err
	}
	result.Data.FillMenuAction(menuActionResult.Data.ToMenuIDMap())
	return result, nil
}

func (a *MenuSrv) Get(ctx context.Context, id uint64, opts ...schema.MenuQueryOptions) (*schema.Menu, error) {
	item, err := a.MenuRepo.Get(ctx, id, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	actions, err := a.QueryActions(ctx, id)
	if err != nil {
		return nil, err
	}
	item.Actions = actions

	return item, nil
}

func (a *MenuSrv) QueryActions(ctx context.Context, id uint64) (schema.MenuActions, error) {
	result, err := a.MenuActionRepo.Query(ctx, schema.MenuActionQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	} else if len(result.Data) == 0 {
		return nil, nil
	}

	resourceResult, err := a.MenuActionResourceRepo.Query(ctx, schema.MenuActionResourceQueryParam{
		MenuID: id,
	})
	if err != nil {
		return nil, err
	}

	result.Data.FillResources(resourceResult.Data.ToActionIDMap())

	return result.Data, nil
}

func (a *MenuSrv) checkName(ctx context.Context, item schema.Menu) error {
	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{
			OnlyCount: true,
		},
		ParentID: &item.ParentID,
		Name:     item.Name,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("名称不能重复")
	}
	return nil
}

func (a *MenuSrv) Create(ctx context.Context, item schema.Menu) (*schema.IDResult, error) {
	if err := a.checkName(ctx, item); err != nil {
		return nil, err
	}

	parentPath, err := a.getParentPath(ctx, item.ParentID)
	if err != nil {
		return nil, err
	}
	item.ParentPath = parentPath
	item.ID = snowflake.MustID()

	err = a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err := a.createActions(ctx, item.ID, item.Actions)
		if err != nil {
			return err
		}

		return a.MenuRepo.Create(ctx, item)
	})
	if err != nil {
		return nil, err
	}

	return schema.NewIDResult(item.ID), nil
}

func (a *MenuSrv) createActions(ctx context.Context, menuID uint64, items schema.MenuActions) error {
	for _, item := range items {
		item.ID = snowflake.MustID()
		item.MenuID = menuID
		err := a.MenuActionRepo.Create(ctx, *item)
		if err != nil {
			return err
		}

		for _, ritem := range item.Resources {
			ritem.ID = snowflake.MustID()
			ritem.ActionID = item.ID
			err := a.MenuActionResourceRepo.Create(ctx, *ritem)
			if err != nil {
				return err
			}
		}

	}
	return nil
}

func (a *MenuSrv) getParentPath(ctx context.Context, parentID uint64) (string, error) {
	if parentID == 0 {
		return "", nil
	}

	pitem, err := a.MenuRepo.Get(ctx, parentID)
	if err != nil {
		return "", err
	} else if pitem == nil {
		return "", errors.ErrInvalidParent
	}

	return a.joinParentPath(pitem.ParentPath, pitem.ID), nil
}

func (a *MenuSrv) joinParentPath(parent string, id uint64) string {
	if parent != "" {
		parent += "/"
	}

	return fmt.Sprintf("%s%d", parent, id)
}

func (a *MenuSrv) Update(ctx context.Context, id uint64, item schema.Menu) error {
	if id == item.ParentID {
		return errors.ErrInvalidParent
	}

	oldItem, err := a.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Name != item.Name {
		if err := a.checkName(ctx, item); err != nil {
			return err
		}
	}

	item.ID = oldItem.ID
	item.Creator = oldItem.Creator
	item.CreatedAt = oldItem.CreatedAt

	if oldItem.ParentID != item.ParentID {
		parentPath, err := a.getParentPath(ctx, item.ParentID)
		if err != nil {
			return err
		}
		item.ParentPath = parentPath
	} else {
		item.ParentPath = oldItem.ParentPath
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err := a.updateActions(ctx, id, oldItem.Actions, item.Actions)
		if err != nil {
			return err
		}

		err = a.updateChildParentPath(ctx, *oldItem, item)
		if err != nil {
			return err
		}

		return a.MenuRepo.Update(ctx, id, item)
	})
}

func (a *MenuSrv) updateActions(ctx context.Context, menuID uint64, oldItems, newItems schema.MenuActions) error {
	addActions, delActions, updateActions := a.compareActions(ctx, oldItems, newItems)

	err := a.createActions(ctx, menuID, addActions)
	if err != nil {
		return err
	}

	for _, item := range delActions {
		err := a.MenuActionRepo.Delete(ctx, item.ID)
		if err != nil {
			return err
		}

		err = a.MenuActionResourceRepo.DeleteByActionID(ctx, item.ID)
		if err != nil {
			return err
		}
	}

	mOldItems := oldItems.ToMap()
	for _, item := range updateActions {
		oitem := mOldItems[item.Code]
		if item.Name != oitem.Name {
			oitem.Name = item.Name
			err := a.MenuActionRepo.Update(ctx, item.ID, *oitem)
			if err != nil {
				return err
			}
		}

		addResources, delResources := a.compareResources(ctx, oitem.Resources, item.Resources)
		for _, aritem := range addResources {
			aritem.ID = snowflake.MustID()
			aritem.ActionID = oitem.ID
			err := a.MenuActionResourceRepo.Create(ctx, *aritem)
			if err != nil {
				return err
			}
		}

		for _, ditem := range delResources {
			err := a.MenuActionResourceRepo.Delete(ctx, ditem.ID)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *MenuSrv) compareActions(ctx context.Context, oldActions, newActions schema.MenuActions) (addList, delList, updateList schema.MenuActions) {
	mOldActions := oldActions.ToMap()
	mNewActions := newActions.ToMap()

	for k, item := range mNewActions {
		if _, ok := mOldActions[k]; ok {
			updateList = append(updateList, item)
			delete(mOldActions, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldActions {
		delList = append(delList, item)
	}
	return
}

func (a *MenuSrv) compareResources(ctx context.Context, oldResources, newResources schema.MenuActionResources) (addList, delList schema.MenuActionResources) {
	mOldResources := oldResources.ToMap()
	mNewResources := newResources.ToMap()

	for k, item := range mNewResources {
		if _, ok := mOldResources[k]; ok {
			delete(mOldResources, k)
			continue
		}
		addList = append(addList, item)
	}

	for _, item := range mOldResources {
		delList = append(delList, item)
	}
	return
}

func (a *MenuSrv) updateChildParentPath(ctx context.Context, oldItem, newItem schema.Menu) error {
	if oldItem.ParentID == newItem.ParentID {
		return nil
	}

	opath := a.joinParentPath(oldItem.ParentPath, oldItem.ID)
	result, err := a.MenuRepo.Query(contextx.NewNoTrans(ctx), schema.MenuQueryParam{
		PrefixParentPath: opath,
	})
	if err != nil {
		return err
	}

	npath := a.joinParentPath(newItem.ParentPath, newItem.ID)
	for _, menu := range result.Data {
		err = a.MenuRepo.UpdateParentPath(ctx, menu.ID, npath+menu.ParentPath[len(opath):])
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *MenuSrv) Delete(ctx context.Context, id uint64) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	result, err := a.MenuRepo.Query(ctx, schema.MenuQueryParam{
		PaginationParam: schema.PaginationParam{OnlyCount: true},
		ParentID:        &id,
	})
	if err != nil {
		return err
	} else if result.PageResult.Total > 0 {
		return errors.New400Response("forbid delete")
	}

	return a.TransRepo.Exec(ctx, func(ctx context.Context) error {
		err = a.MenuActionResourceRepo.DeleteByMenuID(ctx, id)
		if err != nil {
			return err
		}

		err := a.MenuActionRepo.DeleteByMenuID(ctx, id)
		if err != nil {
			return err
		}

		return a.MenuRepo.Delete(ctx, id)
	})
}

func (a *MenuSrv) UpdateStatus(ctx context.Context, id uint64, status int) error {
	oldItem, err := a.MenuRepo.Get(ctx, id)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	} else if oldItem.Status == status {
		return nil
	}

	return a.MenuRepo.UpdateStatus(ctx, id, status)
}
