package service

import (
	"github.com/spark8899/ops-manager/server/service/example"
	"github.com/spark8899/ops-manager/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
