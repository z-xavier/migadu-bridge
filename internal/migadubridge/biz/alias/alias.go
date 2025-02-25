package alias

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/config"

	migadugo "github.com/ZhangXavier/migadu-go"
)

type AliasBiz interface {
	List(*gin.Context) (any, error)
}

type aliasBiz struct {
	ds store.IStore
}

func New(ds store.IStore) AliasBiz {
	return &aliasBiz{ds: ds}
}

func (a *aliasBiz) List(ctx *gin.Context) (any, error) {
	client, err := migadugo.New(
		config.GetConfig().MigaduConf.Email,
		config.GetConfig().MigaduConf.APIKey,
		config.GetConfig().MigaduConf.Domain,
	)
	if err != nil {
		return nil, err
	}
	return client.ListAliases(ctx)
}
