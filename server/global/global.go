package global

import (
	"sync"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spark8899/ops-manager/server/utils/timer"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"github.com/spark8899/ops-manager/server/config"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	OPM_DB     *gorm.DB
	OPM_DBList map[string]*gorm.DB
	OPM_REDIS  *redis.Client
	OPM_CONFIG config.Server
	OPM_VP     *viper.Viper
	// OPM_LOG    *oplogging.Logger
	OPM_LOG                 *zap.Logger
	OPM_Timer               timer.Timer = timer.NewTimerTask()
	OPM_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache
	lock       sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return OPM_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := OPM_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
