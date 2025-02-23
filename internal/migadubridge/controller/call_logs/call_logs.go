package call_logs

import (
	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
)

// CallLogsController 定义了 controller 层需要实现的方法.
type CallLogsController struct {
	b biz.IBiz
}

func New(ds store.IStore) *CallLogsController {
	return &CallLogsController{
		b: biz.NewBiz(ds),
	}
}
