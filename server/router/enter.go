package router

import (
	"github.com/spark8899/ops-manager/server/router/example"
	"github.com/spark8899/ops-manager/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
