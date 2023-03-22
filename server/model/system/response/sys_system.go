package response

import "github.com/spark8899/ops-manager/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
