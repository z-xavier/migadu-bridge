package alias

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/config"
	"migadu-bridge/internal/pkg/model"
	v1 "migadu-bridge/pkg/api/manage/v1"

	migadugo "github.com/ZhangXavier/migadu-go"
)

type AliasBiz interface {
	List(*gin.Context, *v1.ListAliasReq) (*v1.ListAliasResp, error)
}

type aliasBiz struct {
	ds store.IStore
}

func New(ds store.IStore) AliasBiz {
	return &aliasBiz{ds: ds}
}

func (a *aliasBiz) List(ctx *gin.Context, req *v1.ListAliasReq) (*v1.ListAliasResp, error) {
	client, err := migadugo.New(
		config.GetConfig().MigaduConf.Email,
		config.GetConfig().MigaduConf.APIKey,
		config.GetConfig().MigaduConf.Domain,
	)
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
					Id:          id,
					TargetEmail: destination,
					Alias:       alias.Address,
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
