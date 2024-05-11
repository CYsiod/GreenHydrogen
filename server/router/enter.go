package router

import (
	"GreenHydrogen/server/router/equipment"
	"GreenHydrogen/server/router/example"
	"GreenHydrogen/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Equipment equipment.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
