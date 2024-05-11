package request

import (
	"GreenHydrogen/server/model/common/request"
	"GreenHydrogen/server/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
