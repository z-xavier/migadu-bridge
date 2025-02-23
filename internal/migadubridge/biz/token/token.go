package token

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/pkg/api/enum"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type TokenBiz interface {
	Create(*gin.Context, *v1.CreateTokenReq) (*v1.Token, error)
	Delete(*gin.Context, string) error
	Put(*gin.Context, string, *v1.PutTokenReq) (*v1.Token, error)
	Patch(*gin.Context, string, *v1.PatchTokenReq) (*v1.Token, error)
	List(*gin.Context, *v1.ListTokenReq) (*v1.ListTokenResp, error)
	Get(*gin.Context, string) (*v1.Token, error)
}

type tokenBiz struct {
	ds store.IStore
}

func New(ds store.IStore) TokenBiz {
	return &tokenBiz{ds: ds}
}

func (t *tokenBiz) genToken() string {
	return xid.New().String() + xid.New().String()
}

func (t *tokenBiz) Create(ctx *gin.Context, createToken *v1.CreateTokenReq) (*v1.Token, error) {
	token := t.genToken()

	id, err := t.ds.Token().Create(ctx, &model.Token{
		TargetEmail:  createToken.TargetEmail,
		MockProvider: createToken.MockProvider,
		Description:  createToken.Description,
		Token:        token,
		ExpiryAt:     time.Unix(createToken.ExpiryAt, 0),
		LastCalledAt: time.Unix(0, 0),
		Status:       enum.TokenStatusInactive,
	})
	if err != nil {
		return nil, err
	}

	tmp, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return t.transModelToVo(tmp), nil
}

func (t *tokenBiz) Delete(ctx *gin.Context, id string) error {
	return t.ds.Token().DeleteById(ctx, id)
}

func (t *tokenBiz) Put(ctx *gin.Context, id string, req *v1.PutTokenReq) (*v1.Token, error) {
	if _, err := t.ds.Token().GetById(ctx, id); err != nil {
		return nil, err
	}

	updates := map[string]any{
		"description": req.Description,
		"expiry_at":   time.Unix(req.ExpiryAt, 0),
	}

	if err := t.ds.Token().UpdateById(ctx, id, updates); err != nil {
		return nil, err
	}

	tmp, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return t.transModelToVo(tmp), nil
}

func (t *tokenBiz) Patch(ctx *gin.Context, id string, req *v1.PatchTokenReq) (*v1.Token, error) {
	oldToken, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if oldToken.Status == enum.TokenStatusInactive {
		return nil, errmsg.ErrInvalidParameter.SetMessage("token is Inactive")
	}

	updates := map[string]any{
		"status": req.Status,
	}
	if err = t.ds.Token().UpdateById(ctx, id, updates); err != nil {
		return nil, err
	}
	tmp, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return t.transModelToVo(tmp), nil
}

func (t *tokenBiz) List(ctx *gin.Context, listTokenReq *v1.ListTokenReq) (*v1.ListTokenResp, error) {
	if listTokenReq.Page == 0 {
		listTokenReq.Page = 1
	}

	if listTokenReq.PageSize == 0 {
		listTokenReq.PageSize = 10
	}

	if len(listTokenReq.OrderBy) == 0 {
		listTokenReq.OrderBy = []string{"updated_at:desc"}
	}

	cond := map[string][]any{}
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
	if listTokenReq.UpdatedAtBegin != 0 {
		cond["updated_at >= ?"] = []any{time.Unix(listTokenReq.UpdatedAtBegin, 0)}
	}
	if listTokenReq.UpdatedAtEnd != 0 {
		cond["updated_at <= ?"] = []any{time.Unix(listTokenReq.UpdatedAtEnd, 0)}
	}
	if listTokenReq.Status != 0 {
		cond["status = ?"] = []any{listTokenReq.Status}
	}

	var orderBy []any
	for _, o := range listTokenReq.OrderBy {
		parts := strings.Split(o, ":")
		if len(parts) == 2 {
			orderBy = append(orderBy, fmt.Sprintf("%s %s", parts[0], parts[1]))
		}
	}

	count, tokenList, err := t.ds.Token().List(ctx, listTokenReq.Page, listTokenReq.PageSize, cond, orderBy)
	if err != nil {
		return nil, err
	}

	tokens := make([]*v1.Token, 0, len(tokenList))
	for _, tmp := range tokenList {
		tokens = append(tokens, t.transModelToVo(tmp))
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

func (t *tokenBiz) Get(ctx *gin.Context, id string) (*v1.Token, error) {
	tmp, err := t.ds.Token().GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return t.transModelToVo(tmp), nil
}

func (t *tokenBiz) transModelToVo(token *model.Token) *v1.Token {
	return &v1.Token{
		Id:           token.Id,
		TargetEmail:  token.TargetEmail,
		MockProvider: token.MockProvider,
		Description:  token.Description,
		Token:        token.Token,
		ExpiryAt:     token.ExpiryAt.Unix(),
		LastCalledAt: token.LastCalledAt.Unix(),
		Status:       token.Status,
		CreatedAt:    token.CreatedAt.Unix(),
		UpdatedAt:    token.UpdatedAt.Unix(),
	}
}
