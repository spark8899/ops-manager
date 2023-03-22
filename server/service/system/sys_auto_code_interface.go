package system

import (
	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/model/system/response"
)

type Database interface {
	GetDB(businessDB string) (data []response.Db, err error)
	GetTables(businessDB string, dbName string) (data []response.Table, err error)
	GetColumn(businessDB string, tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database(businessDB string) Database {

	if businessDB == "" {
		switch global.OPM_CONFIG.System.DbType {
		case "mysql":
			return AutoCodeMysql
		case "pgsql":
			return AutoCodePgsql
		default:
			return AutoCodeMysql
		}
	} else {
		for _, info := range global.OPM_CONFIG.DBList {
			if info.AliasName == businessDB {

				switch info.Type {
				case "mysql":
					return AutoCodeMysql
				case "mssql":
					return AutoCodeMssql
				case "pgsql":
					return AutoCodePgsql
				case "oracle":
					return AutoCodeOracle
				default:
					return AutoCodeMysql
				}
			}
		}
		return AutoCodeMysql
	}

}
