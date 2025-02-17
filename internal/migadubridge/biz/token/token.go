package token

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/pkg/api/enum"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type TokenBiz interface {
	Create(*gin.Context, *v1.CreateTokenReq) (*v1.CreateTokenResp, error)
	Delete(*gin.Context, string) error
	Update(*gin.Context, string, *v1.UpdateTokenReq) error
	List(*gin.Context, *v1.ListTokenReq) (*v1.ListTokenResp, error)
	Get(*gin.Context, string) (*v1.GetTokenResp, error)
}

type tokenBiz struct {
	ds store.IStore
}

func New(ds store.IStore) TokenBiz {
	return &tokenBiz{ds: ds}
}

func (t *tokenBiz) Create(ctx *gin.Context, createToken *v1.CreateTokenReq) (*v1.CreateTokenResp, error) {
	token := xid.New().String()
	id, err := t.ds.Token().Create(ctx, &model.Token{
		TargetEmail:  createToken.TargetEmail,
		MockProvider: createToken.MockProvider,
		Description:  createToken.Description,
		Token:        token,
		ExpiryAt:     time.Unix(createToken.ExpiryAt, 0),
		Status:       enum.TokenStatusInactive,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateTokenResp{
		Id:    id,
		Token: token,
	}, nil
}

func (t *tokenBiz) Delete(ctx *gin.Context, id string) error {
	return t.ds.Token().DeleteById(ctx, id)
}

func (t *tokenBiz) Update(ctx *gin.Context, id string, req *v1.UpdateTokenReq) error {
	_, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return err
	}
	updates := map[string]any{
		"description": req.Description,
	}
	if req.Status != 0 {
		updates["status"] = req.Status
	}
	if req.ExpiryAt != 0 {
		updates["expiry_at"] = time.Unix(req.ExpiryAt, 0)
	}
	return t.ds.Token().UpdateById(ctx, id, updates)
}

func (t *tokenBiz) List(ctx *gin.Context, listTokenReq *v1.ListTokenReq) (*v1.ListTokenResp, error) {
	if listTokenReq.Page == 0 {
		listTokenReq.Page = 1
	}

	if listTokenReq.PageSize == 0 {
		listTokenReq.PageSize = 10
	}

	cond := map[string][]any{}
	if listTokenReq.Id != "" {
		cond["id = ?"] = []any{listTokenReq.Id}
	}
	if listTokenReq.TargetEmail != "" {
		cond["targetEmail like ?"] = []any{"%" + listTokenReq.TargetEmail + "%"}
	}
	if listTokenReq.MockProvider != "" {
		cond["mockProvider = ?"] = []any{listTokenReq.MockProvider}
	}
	if listTokenReq.Description != "" {
		cond["description like ?"] = []any{"%" + listTokenReq.Description + "%"}
	}
	if listTokenReq.ExpiryAtBegin != 0 {
		cond["expiry_at >= ?"] = []any{time.Unix(listTokenReq.ExpiryAtBegin, 0)}
	}
	if listTokenReq.ExpiryAtEnd != 0 {
		cond["expiry_at <= ?"] = []any{time.Unix(listTokenReq.ExpiryAtEnd, 0)}
	}
	if listTokenReq.LastCalledAtBegin != 0 {
		cond["last_called_at >= ?"] = []any{time.Unix(listTokenReq.LastCalledAtBegin, 0)}
	}
	if listTokenReq.LastCalledAtEnd != 0 {
		cond["last_called_at <= ?"] = []any{time.Unix(listTokenReq.LastCalledAtEnd, 0)}
	}
	if listTokenReq.Status != 0 {
		cond["status = ?"] = []any{listTokenReq.Status}
	}

	count, tokenList, err := t.ds.Token().List(ctx, listTokenReq.Page, listTokenReq.PageSize, cond)
	if err != nil {
		return nil, err
	}

	tokens := make([]*v1.Token, 0, len(tokenList))
	for _, tmp := range tokenList {
		tokens = append(tokens, &v1.Token{
			Id:           tmp.Id,
			TargetEmail:  tmp.TargetEmail,
			MockProvider: tmp.MockProvider,
			Description:  tmp.Description,
			Token:        tmp.Token,
			ExpiryAt:     tmp.ExpiryAt.Unix(),
			LastCalledAt: tmp.LastCalledAt.Unix(),
			Status:       tmp.Status,
			CreatedAt:    tmp.CreatedAt.Unix(),
			UpdatedAt:    tmp.UpdatedAt.Unix(),
		})
	}

	return &v1.ListTokenResp{
		Page: v1.Page{
			Page:     listTokenReq.Page,
			PageSize: listTokenReq.PageSize,
			Total:    uint64(count),
		},
		List: tokens,
	}, nil
}

func (t *tokenBiz) Get(ctx *gin.Context, id string) (*v1.GetTokenResp, error) {
	tmp, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &v1.GetTokenResp{
		Token: v1.Token{
			Id:           tmp.Id,
			TargetEmail:  tmp.TargetEmail,
			MockProvider: tmp.MockProvider,
			Description:  tmp.Description,
			Token:        tmp.Token,
			ExpiryAt:     tmp.ExpiryAt.Unix(),
			LastCalledAt: tmp.LastCalledAt.Unix(),
			Status:       tmp.Status,
			CreatedAt:    tmp.CreatedAt.Unix(),
			UpdatedAt:    tmp.UpdatedAt.Unix(),
		},
	}, nil
}
