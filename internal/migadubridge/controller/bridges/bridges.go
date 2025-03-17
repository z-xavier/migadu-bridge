package bridges

import (
	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
)

func New(ds store.IStore) *BridgesController {
	return &BridgesController{
		b: biz.NewBiz(ds),
	}
}

// BridgesController 定义了 controller 层需要实现的方法.
type BridgesController struct {
	b biz.IBiz
}
