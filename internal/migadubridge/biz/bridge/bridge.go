package bridge

import "context"

type BridgeBiz interface {
	AliasRandomNew(ctx context.Context, domain, desc string) error
}

// New 创建一个实现了 UserBiz 接口的实例.
func New() BridgeBiz {
	return &bridgeBiz{}
}
