package response

import (
	"github.com/spark8899/ops-manager/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
