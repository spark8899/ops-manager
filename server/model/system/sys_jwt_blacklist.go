package system

import (
	"github.com/spark8899/ops-manager/server/global"
)

type JwtBlacklist struct {
	global.OPM_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
