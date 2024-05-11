package request

import (
	"github.com/CYsiod/GreenHydrogen/server/model/common/request"
	"github.com/CYsiod/GreenHydrogen/server/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
