package service

import (
	"GreenHydrogen/server/service/equipment"
	"GreenHydrogen/server/service/example"
	"GreenHydrogen/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup    system.ServiceGroup
	ExampleServiceGroup   example.ServiceGroup
	EquipmentServiceGroup equipment.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
