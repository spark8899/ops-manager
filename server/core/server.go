package core

import (
	"fmt"
	"time"

	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/initialize"
	"github.com/spark8899/ops-manager/server/service/system"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.OPM_CONFIG.System.UseMultipoint || global.OPM_CONFIG.System.UseRedis {
		// init redis service
		initialize.Redis()
	}

	// from db load jwt data
	if global.OPM_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.OPM_CONFIG.System.Addr)
	s := initServer(address, Router)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.OPM_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	ops-manager is running.
`, address)
	global.OPM_LOG.Error(s.ListenAndServe().Error())
}
