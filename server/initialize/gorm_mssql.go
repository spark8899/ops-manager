/*
 * @Author: spark8899
 * @Date: 2022-12-08 17:25:49
 * @LastEditors: spark8899
 * @LastEditTime: 2022-12-08 18:00:00
 * @FilePath: \server\initialize\gorm_mssql.go
 * @Description: default setting `customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"github.com/spark8899/ops-manager/server/config"
	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/initialize/internal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql init Mssql
// Author spark8899
func GormMssql() *gorm.DB {
	m := global.OPM_CONFIG.Mssql
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string length
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMssqlByConfig init Mysql
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.Dbname == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(), // DSN data source name
		DefaultStringSize: 191,     // string length
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE=InnoDB")
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}
