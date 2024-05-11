package v1

import (
	"github.com/CYsiod/GreenHydrogen/server/api/v1/equipment"
	"github.com/CYsiod/GreenHydrogen/server/api/v1/example"
	"github.com/CYsiod/GreenHydrogen/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	EquipmentApiGroup equipment.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
