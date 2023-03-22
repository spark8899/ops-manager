package request

import (
	"github.com/spark8899/ops-manager/server/model/common/request"
	"github.com/spark8899/ops-manager/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
