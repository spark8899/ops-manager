package internal

import (
	"fmt"

	"github.com/spark8899/ops-manager/server/global"
	"gorm.io/gorm/logger"
)

type writer struct {
	logger.Writer
}

// NewWriter writer build function
// Author spark8899
func NewWriter(w logger.Writer) *writer {
	return &writer{Writer: w}
}

// Printf format print log
// Author spark8899
func (w *writer) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.OPM_CONFIG.System.DbType {
	case "mysql":
		logZap = global.OPM_CONFIG.Mysql.LogZap
	case "pgsql":
		logZap = global.OPM_CONFIG.Pgsql.LogZap
	}
	if logZap {
		global.OPM_LOG.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
