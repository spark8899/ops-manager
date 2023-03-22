package initialize

import (
	_ "github.com/spark8899/ops-manager/server/source/example"
	_ "github.com/spark8899/ops-manager/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
