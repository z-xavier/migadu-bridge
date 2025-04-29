package provider

import (
	"context"
	"errors"

	"migadu-bridge/pkg/api/enum"
)

type Biz interface {
	AliasRandomNew(ctx context.Context, domain, desc string) error
}

// New 创建一个实现了 Biz 接口的实例.
func New(providerEnum enum.ProviderEnum) (Biz, error) {
	switch providerEnum {
	case enum.ProviderEnumAddy:
		return NewAddy(), nil
	case enum.ProviderEnumSimpleLogin:
		return NewSimpleLogin(), nil
	default:
		return nil, errors.New("unknown provider")
	}
}
