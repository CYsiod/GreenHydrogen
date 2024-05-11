package v1

import (
	"GreenHydrogen/server/api/v1/equipment"
	"GreenHydrogen/server/api/v1/example"
	"GreenHydrogen/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	EquipmentApiGroup equipment.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
