package request

import (
	"GreenHydrogen/server/model/common/request"
	"GreenHydrogen/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
