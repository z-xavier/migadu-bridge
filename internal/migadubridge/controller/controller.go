package controller

import "migadu-bridge/internal/migadubridge/biz"

type Controller struct {
	b biz.IBiz
}

// New 创建一个 controller.
func New() *Controller {
	return &Controller{b: biz.NewBiz()}
}
