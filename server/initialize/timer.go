package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/spark8899/ops-manager/server/config"
	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/utils"
)

func Timer() {
	if global.OPM_CONFIG.Timer.Start {
		for i := range global.OPM_CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.OPM_CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.OPM_Timer.AddTaskByFunc("ClearDB", global.OPM_CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.OPM_DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.OPM_CONFIG.Timer.Detail[i])
		}
	}
}
