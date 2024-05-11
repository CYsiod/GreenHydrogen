package router

import (
	"github.com/CYsiod/GreenHydrogen/server/router/equipment"
	"github.com/CYsiod/GreenHydrogen/server/router/example"
	"github.com/CYsiod/GreenHydrogen/server/router/system"
)

type RouterGroup struct {
	System    system.RouterGroup
	Example   example.RouterGroup
	Equipment equipment.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
