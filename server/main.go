package main

import (
	"go.uber.org/zap"

	"github.com/spark8899/ops-manager/server/core"
	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/initialize"
)

//go:generate go mod tidy
//go:generate go mod download

// @title                       Swagger Example API
// @version                     0.0.1
// @description                 This is a sample Server pets
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        x-token
// @BasePath                    /
func main() {
	global.OPM_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.OPM_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.OPM_LOG)
	global.OPM_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	initialize.DBList()
	if global.OPM_DB != nil {
		//initialize.RegisterTables() // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.OPM_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
