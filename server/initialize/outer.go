package initialize

import (
	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/spark8899/ops-manager/server/global"
	"github.com/spark8899/ops-manager/server/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.OPM_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.OPM_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}

	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
}
