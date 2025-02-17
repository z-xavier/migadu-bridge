package provider

import (
	"context"
	"errors"
	"sync"

	"github.com/ZhangXavier/migadu-go"

	"migadu-bridge/internal/pkg/config"
	"migadu-bridge/pkg/api/enum"
)

type ProviderBiz interface {
	AliasRandomNew(ctx context.Context, domain, desc string) error
}

// New 创建一个实现了 ProviderBiz 接口的实例.
func New(providerEnum enum.ProviderEnum) (ProviderBiz, error) {
	switch providerEnum {
	case enum.ProviderEnumAddy:
		return NewAddy(), nil
	case enum.ProviderEnumSimpleLogin:
		return NewSimpleLogin(), nil
	default:
		return nil, errors.New("unknown provider")
	}
}

var MigaduClient = sync.OnceValues(func() (*migadu.Client, error) {
	return migadu.New(
		config.GetConfig().MigaduConf.Email,
		config.GetConfig().MigaduConf.APIKey,
		config.GetConfig().MigaduConf.Domain)
})
