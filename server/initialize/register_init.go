package initialize

import (
	_ "GreenHydrogen/server/source/example"
	_ "GreenHydrogen/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
