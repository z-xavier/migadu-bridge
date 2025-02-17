package token

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type TokenBiz interface {
	Create(*gin.Context, *v1.CreateTokenReq) (*v1.CreateTokenResp, error)
}

type tokenBiz struct {
	ds store.IStore
}

func New(ds store.IStore) TokenBiz {
	return &tokenBiz{ds: ds}
}

func (t *tokenBiz) Create(ctx *gin.Context, createToken *v1.CreateTokenReq) (*v1.CreateTokenResp, error) {
	//t.ds.DB().Token.Create().
	//	SetID().
	//	SetCreatedAt(time.Now()).
	//	Save(ctx)

	return nil, nil
}
