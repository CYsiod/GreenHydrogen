package request

import (
	"github.com/CYsiod/GreenHydrogen/server/model/common/request"
	"github.com/CYsiod/GreenHydrogen/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
