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
	Token() token.TokenBiz
	// CallLog 调用记录
	CallLog() call_log.CallLogBiz
	// Alias 当前别名查看
	Alias() alias.AliasBiz
	// Bridge 转发外部请求
	Bridge() bridge.BridgeBiz
}

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(ds store.IStore) IBiz {
	return &biz{ds: ds}
}

func (b *biz) Token() token.TokenBiz {
	return token.New(b.ds)
}

func (b *biz) CallLog() call_log.CallLogBiz {
	return call_log.New(b.ds)
}

func (b *biz) Alias() alias.AliasBiz {
	return alias.New(b.ds)
}

func (b *biz) Bridge() bridge.BridgeBiz {
	return bridge.New(b.ds)
}
