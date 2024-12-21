package biz

import "migadu-bridge/internal/migadubridge/biz/bridge"

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	Bridge() bridge.BridgeBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct{}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz() *biz {
	return &biz{}
}

func (b *biz) Bridge() bridge.BridgeBiz {
	return bridge.New()
}
