package token

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"gorm.io/gorm"

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
	if _, err := t.ds.Token().GetById(ctx, id); err != nil {
		return err
	}
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

func (t *tokenBiz) List(ctx *gin.Context, req *v1.ListTokenReq) (*v1.ListTokenResp, error) {
	if req.Page == 0 {
		req.Page = 1
	}

	if req.PageSize == 0 {
		req.PageSize = 10
	}

	if len(req.OrderBy) == 0 {
		req.OrderBy = []string{"updated_at:desc"}
	}

	cond := map[string][]any{}
	if req.TargetEmail != "" {
		cond["targetEmail like ?"] = []any{"%" + req.TargetEmail + "%"}
	}
	if req.MockProvider != "" {
		cond["mockProvider = ?"] = []any{req.MockProvider}
	}
	if req.Description != "" {
		cond["description like ?"] = []any{"%" + req.Description + "%"}
	}
	if req.ExpiryAtBegin != 0 {
		cond["expiry_at >= ?"] = []any{time.Unix(req.ExpiryAtBegin, 0)}
	}
	if req.ExpiryAtEnd != 0 {
		cond["expiry_at <= ?"] = []any{time.Unix(req.ExpiryAtEnd, 0)}
	}
	if req.LastCalledAtBegin != 0 {
		cond["last_called_at >= ?"] = []any{time.Unix(req.LastCalledAtBegin, 0)}
	}
	if req.LastCalledAtEnd != 0 {
		cond["last_called_at <= ?"] = []any{time.Unix(req.LastCalledAtEnd, 0)}
	}
	if req.UpdatedAtBegin != 0 {
		cond["updated_at >= ?"] = []any{time.Unix(req.UpdatedAtBegin, 0)}
	}
	if req.UpdatedAtEnd != 0 {
		cond["updated_at <= ?"] = []any{time.Unix(req.UpdatedAtEnd, 0)}
	}
	if req.Status != 0 {
		cond["status = ?"] = []any{req.Status}
	}

	var orderBy []any
	for _, o := range req.OrderBy {
		parts := strings.Split(o, ":")
		if len(parts) == 2 {
			orderBy = append(orderBy, fmt.Sprintf("%s %s", parts[0], parts[1]))
		}
	}

	count, tokenList, err := t.ds.Token().ListWithPage(ctx, req.Page, req.PageSize, cond, orderBy)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	tokens := make([]*v1.Token, 0, len(tokenList))
	for _, tmp := range tokenList {
		tokens = append(tokens, t.transModelToVo(tmp))
	}

	return &v1.ListTokenResp{
		Page: v1.Page{
			Page:     req.Page,
			PageSize: req.PageSize,
			Total:    count,
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
