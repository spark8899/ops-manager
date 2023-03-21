//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package app

import (
	"github.com/google/wire"

	"github.com/spark8899/ops-manager/server/internal/app/api"
	"github.com/spark8899/ops-manager/server/internal/app/dao"
	"github.com/spark8899/ops-manager/server/internal/app/module/adapter"
	"github.com/spark8899/ops-manager/server/internal/app/router"
	"github.com/spark8899/ops-manager/server/internal/app/service"
)

func BuildInjector() (*Injector, func(), error) {
	wire.Build(
		InitGormDB,
		dao.RepoSet,
		InitAuth,
		InitCasbin,
		InitGinEngine,
		service.ServiceSet,
		api.APISet,
		router.RouterSet,
		adapter.CasbinAdapterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
