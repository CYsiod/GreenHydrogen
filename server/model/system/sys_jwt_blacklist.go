package system

import (
	"github.com/CYsiod/GreenHydrogen/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
