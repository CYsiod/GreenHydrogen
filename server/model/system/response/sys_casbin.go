package response

import (
	"github.com/CYsiod/GreenHydrogen/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
