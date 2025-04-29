package bridges

import (
	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
)

func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

// Controller 定义了 controller 层需要实现的方法.
type Controller struct {
	b biz.IBiz
}
