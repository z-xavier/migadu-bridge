package alias

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/migadu"
	"migadu-bridge/internal/pkg/model"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type Biz interface {
	List(*gin.Context, *v1.ListAliasReq) (*v1.ListAliasResp, error)
	Delete(*gin.Context, string) error
}

type aliasBiz struct {
	ds store.IStore
}

func New(ds store.IStore) Biz {
	return &aliasBiz{ds: ds}
}

func (a *aliasBiz) List(ctx *gin.Context, req *v1.ListAliasReq) (*v1.ListAliasResp, error) {
	client, err := migadu.MigaduClient()
	if err != nil {
		return nil, err
	}

	aliases, err := client.ListAliases(ctx)
	if err != nil {
		return nil, err
	}

	total, id := int64(0), int64(0)
	aliasList := make([]*v1.Alias, 0)

	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize

	targetEmailList := make([]string, 0, req.PageSize)

	for _, alias := range aliases {
		for _, destination := range alias.Destinations {
			total++
			if total > start && total <= end {
				id++
				aliasList = append(aliasList, &v1.Alias{
					Id:               id,
					TargetEmail:      destination,
					Alias:            alias.Address,
					Expireable:       alias.Expireable,
					ExpiresOn:        alias.ExpiresOn,
					IsInternal:       alias.IsInternal,
					RemoveUponExpiry: alias.RemoveUponExpiry,
				})
				targetEmailList = append(targetEmailList, destination)
			}
		}
	}

	// 获取匹配的 token
	targetList, err := a.ds.Token().ListByTargetEmail(ctx, targetEmailList)
	if err != nil {
		return nil, err
	}

	targetIdList := make([]string, 0, len(targetList))
	for _, token := range targetList {
		targetIdList = append(targetIdList, token.Id)
	}

	// 获取匹配的 callLog
	callLogList, err := a.ds.CallLog().ListByTokenId(ctx, targetIdList)
	if err != nil {
		return nil, err
	}

	// 构造 targetMap id -> tokenModel
	targetMap := make(map[string]*model.Token)
	for _, token := range targetList {
		targetMap[token.Id] = token
	}

	// 构造 callLogMap alias_targetEmail -> callLogModel
	callLogMap := make(map[string]*model.CallLog)
	for _, log := range callLogList {
		if token, ok := targetMap[log.TokenId]; ok {
			callLogMap[fmt.Sprintf("%s_%s", log.GenAlias, token.TargetEmail)] = log
		}
	}

	for _, alias := range aliasList {
		if log, ok := callLogMap[fmt.Sprintf("%s_%s", alias.Alias, alias.TargetEmail)]; ok {
			alias.CallLogId = log.Id
			alias.TokenId = log.TokenId
			if token, ok := targetMap[log.TokenId]; ok {
				alias.MockProvider = token.MockProvider
			}
			alias.Description = log.Description
			alias.RequestAt = log.RequestAt.Unix()
		}
	}

	return &v1.ListAliasResp{
		Page: v1.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    total,
		},
		List: aliasList,
	}, nil
}

func (a *aliasBiz) Delete(ctx *gin.Context, alias string) error {
	client, err := migadu.MigaduClient()
	if err != nil {
		return err
	}
	return client.DeleteAlias(ctx, alias)
}
