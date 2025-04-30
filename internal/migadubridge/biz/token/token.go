package token

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/internal/pkg/model"
	"migadu-bridge/pkg/api/enum"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

type Biz interface {
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

func New(ds store.IStore) Biz {
	return &tokenBiz{ds: ds}
}

func (t *tokenBiz) genToken() (string, error) {
	// 生成 32 字节 ( 256 位 ) 的随机数据
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		// 在实际应用中应妥善处理错误
		return "", err
	}
	// 使用URL安全的base64编码
	return base64.URLEncoding.EncodeToString(b), nil
}

func (t *tokenBiz) Create(ctx *gin.Context, createToken *v1.CreateTokenReq) (*v1.Token, error) {
	token, err := t.genToken()
	if err != nil {
		log.C(ctx).WithError(err).Error("gen token error")
		return nil, err
	}

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
		cond["last_called_at IS NOT NULL"] = []any{}
	}
	if req.LastCalledAtEnd != 0 {
		cond["last_called_at <= ?"] = []any{time.Unix(req.LastCalledAtEnd, 0)}
		cond["last_called_at IS NOT NULL"] = []any{}
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
	var lastCalledUnix int64
	if lastCalledAt, _ := token.LastCalledAt.Value(); lastCalledAt != nil {
		if t, ok := lastCalledAt.(time.Time); ok {
			lastCalledAt = t.Unix()
		}
	}

	return &v1.Token{
		Id:           token.Id,
		TargetEmail:  token.TargetEmail,
		MockProvider: token.MockProvider,
		Description:  token.Description,
		Token:        token.Token,
		ExpiryAt:     token.ExpiryAt.Unix(),
		LastCalledAt: lastCalledUnix,
		Status:       token.Status,
		CreatedAt:    token.CreatedAt.Unix(),
		UpdatedAt:    token.UpdatedAt.Unix(),
	}
}
