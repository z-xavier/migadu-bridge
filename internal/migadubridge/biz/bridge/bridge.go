package bridge

import (
	"context"
	"sync"

	"github.com/ZhangXavier/migadu-go"

	"migadu-bridge/internal/pkg/config"
)

type BridgeBiz interface {
	AliasRandomNew(ctx context.Context, domain, desc string) error
}

// New 创建一个实现了 UserBiz 接口的实例.
func New() BridgeBiz {
	return &bridgeBiz{}
}

var MigaduClient = sync.OnceValues(func() (*migadu.Client, error) {
	return migadu.New(
		config.GetConfig().MigaduConf.Email,
		config.GetConfig().MigaduConf.APIKey,
		config.GetConfig().MigaduConf.Domain)
})
