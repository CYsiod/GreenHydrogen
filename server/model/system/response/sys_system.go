package response

import "github.com/CYsiod/GreenHydrogen/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
