package bridge

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/pkg/api/enum"
	"migadu-bridge/pkg/api/manage/bridge/addy"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

type Biz interface {
	SLAliasRandomNew(c *gin.Context, req *sl.AliasRandomNewReq) (*sl.Alias, int, error)
	AddyAliases(c *gin.Context, req *addy.CreateAliasReq) (*addy.CreateAliasResp, int, error)
}

type bridgeBiz struct {
	ds store.IStore
}

func New(ds store.IStore) Biz {
	return &bridgeBiz{ds: ds}
}

func (b *bridgeBiz) checkToken(c *gin.Context, mockProvider enum.ProviderEnum, tokenString string) (*model.Token, error) {
	token, err := b.ds.Token().GetActiveToken(c, mockProvider, tokenString)
	if err != nil {
		log.C(c).WithError(err).Error("Biz CheckToken")
		return nil, err
	}

	if err = b.ds.Token().UpdateById(c, token.Id, map[string]any{
		"last_called_at": c.GetTime(common.XRequestTime),
		"status":         enum.TokenStatusActive,
	}); err != nil {
		log.C(c).WithError(err).Error("Biz UpdateToken")
		return nil, err
	}

	return token, nil
}

func (b *bridgeBiz) log(c *gin.Context, token *model.Token, desc string) (string, error) {
	logId, err := b.ds.CallLog().Create(c, &model.CallLog{
		TokenId:     token.Id,
		Description: desc,
		RequestPath: c.Request.URL.Path,
		RequestAt:   c.GetTime(common.XRequestTime),
	})
	if err != nil {
		log.C(c).WithError(err).Error("Biz CreateCallLog")
		return "", err
	}
	return logId, nil
}

func (b *bridgeBiz) logAlias(c *gin.Context, logId, genAlias string) error {
	if err := b.ds.CallLog().Update(c, logId, &model.CallLog{
		GenAlias: genAlias,
	}); err != nil {
		log.C(c).WithError(err).Error("Biz UpdateCallLog")
		return err
	}
	return nil
}
