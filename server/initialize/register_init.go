package initialize

import (
	_ "github.com/CYsiod/GreenHydrogen/server/source/example"
	_ "github.com/CYsiod/GreenHydrogen/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
