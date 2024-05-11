package request

import (
	"github.com/CYsiod/GreenHydrogen/server/model/common/request"
	"github.com/CYsiod/GreenHydrogen/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
