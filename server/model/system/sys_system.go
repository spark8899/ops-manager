package system

import (
	"github.com/spark8899/ops-manager/server/config"
)

// 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}
