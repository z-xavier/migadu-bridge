package migadu

import (
	"sync"

	"github.com/ZhangXavier/migadu-go"

	"migadu-bridge/internal/pkg/config"
)

var MigaduClient = sync.OnceValues(func() (*migadu.Client, error) {
	return migadu.New(
		config.GetConfig().MigaduConf.Email,
		config.GetConfig().MigaduConf.APIKey,
		config.GetConfig().MigaduConf.Domain)
})
