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

// Controller 定义了 controller 层需要实现的方法.
type Controller struct {
	b biz.IBiz
}

// New 创建一个 token controller.
func New(ds store.IStore) *Controller {
	return &Controller{b: biz.NewBiz(ds)}
}

func (tc *Controller) Create(c *gin.Context) (any, error) {
	log.C(c).Infof("create token begin")

	var r v1.CreateTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("create token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Create(c, &r)
}

func (tc *Controller) Delete(c *gin.Context) (any, error) {
	log.C(c).Infof("delete token begin")

	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return nil, tc.b.Token().Delete(c, tokenId)
}

func (tc *Controller) Put(c *gin.Context) (any, error) {
	log.C(c).Infof("put token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PutTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("put token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Put(c, tokenId, &r)
}

func (tc *Controller) Patch(c *gin.Context) (any, error) {
	log.C(c).Infof("patch token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	var r v1.PatchTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("patch token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().Patch(c, tokenId, &r)
}

func (tc *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Infof("list token begin")

	var r v1.ListTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("list token request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return tc.b.Token().List(c, &r)
}

func (tc *Controller) Get(c *gin.Context) (any, error) {
	log.C(c).Infof("get token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Error("token id is empty")
		return nil, errmsg.ErrBind.SetMessage("token id is required")
	}

	return tc.b.Token().Get(c, tokenId)
}
