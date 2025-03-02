package tokens

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

// TokenController 定义了 controller 层需要实现的方法.
type TokenController struct {
	b biz.IBiz
}

// New 创建一个 token controller.
func New(ds store.IStore) *TokenController {
	return &TokenController{b: biz.NewBiz(ds)}
}

func (tc *TokenController) Create(c *gin.Context) (any, error) {
	log.C(c).Infof("create token begin")

	var r v1.CreateTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("create token request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Create(c, &r)
}

func (tc *TokenController) Delete(c *gin.Context) (any, error) {
	log.C(c).Infof("delete token begin")

	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return nil, tc.b.Token().Delete(c, tokenId)
}

func (tc *TokenController) Put(c *gin.Context) (any, error) {
	log.C(c).Infof("put token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PutTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Infof("put token request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Put(c, tokenId, &r)
}

func (tc *TokenController) Patch(c *gin.Context) (any, error) {
	log.C(c).Infof("patch token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PatchTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Infof("patch token request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Patch(c, tokenId, &r)
}

func (tc *TokenController) List(c *gin.Context) (any, error) {
	log.C(c).Infof("list token begin")

	var r v1.ListTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("list token request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().List(c, &r)
}

func (tc *TokenController) Get(c *gin.Context) (any, error) {
	log.C(c).Infof("get token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return tc.b.Token().Get(c, tokenId)
}
