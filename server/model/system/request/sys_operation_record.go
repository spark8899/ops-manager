package request

import (
	"github.com/spark8899/ops-manager/server/model/common/request"
	"github.com/spark8899/ops-manager/server/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
