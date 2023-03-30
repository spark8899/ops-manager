package core

import (
	"fmt"
	"os"

	"github.com/spark8899/ops-manager/server/core/internal"
	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap get zap.Logger
// Author spark8899
func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.OPM_CONFIG.Zap.Director); !ok { // Determine whether there is a Director folder
		fmt.Printf("create %v directory\n", global.OPM_CONFIG.Zap.Director)
		_ = os.Mkdir(global.OPM_CONFIG.Zap.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger = zap.New(zapcore.NewTee(cores...))

	if global.OPM_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
