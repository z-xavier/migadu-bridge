package biz

import (
	"migadu-bridge/internal/migadubridge/biz/alias"
	"migadu-bridge/internal/migadubridge/biz/bridge"
	"migadu-bridge/internal/migadubridge/biz/call_log"
	"migadu-bridge/internal/migadubridge/biz/token"
	"migadu-bridge/internal/migadubridge/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	// Token 管理 token
	Token() token.Biz
	// CallLog 调用记录
	CallLog() call_log.Biz
	// Alias 当前别名查看
	Alias() alias.Biz
	// Bridge 转发外部请求
	Bridge() bridge.Biz
}

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) IBiz {
	return &biz{ds: ds}
}

func (b *biz) Token() token.Biz {
	return token.New(b.ds)
}

func (b *biz) CallLog() call_log.Biz {
	return call_log.New(b.ds)
}

func (b *biz) Alias() alias.Biz {
	return alias.New(b.ds)
}

func (b *biz) Bridge() bridge.Biz {
	return bridge.New(b.ds)
}
