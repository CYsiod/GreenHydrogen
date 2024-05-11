package response

import "GreenHydrogen/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
