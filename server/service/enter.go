package service

import (
	"github.com/CYsiod/GreenHydrogen/server/service/equipment"
	"github.com/CYsiod/GreenHydrogen/server/service/example"
	"github.com/CYsiod/GreenHydrogen/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	EquipmentServiceGroup equipment.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
