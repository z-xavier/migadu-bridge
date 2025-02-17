package tokens

import (
	"errors"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/common"
	"migadu-bridge/internal/pkg/core"
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

func (tc *TokenController) Create(c *gin.Context) {
	log.C(c).Infof("create token begin")

	var r v1.CreateTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("create token request parse error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	resp, err := tc.b.Token().Create(c, &r)
	if err != nil {
		log.C(c).Errorf("create token error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	log.C(c).Infof("create token end")
}

func (tc *TokenController) Delete(c *gin.Context) {
	log.C(c).Infof("delete token begin")

	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		core.WriteResponse(c, errors.New("token id is required"), nil)
		return
	}

	if err := tc.b.Token().Delete(c, tokenId); err != nil {
		log.C(c).Errorf("delete token error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
	log.C(c).Infof("delete token end")
}

func (tc *TokenController) Update(c *gin.Context) {
	log.C(c).Infof("update token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		core.WriteResponse(c, errors.New("token id is required"), nil)
		return
	}

	var r v1.UpdateTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Infof("update token request parse error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	if err := tc.b.Token().Update(c, tokenId, &r); err != nil {
		log.C(c).Errorf("update token error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, nil)
	log.C(c).Infof("update token end")
}

func (tc *TokenController) List(c *gin.Context) {
	log.C(c).Infof("list token begin")

	var r v1.ListTokenReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("list token request parse error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	resp, err := tc.b.Token().List(c, &r)
	if err != nil {
		log.C(c).Errorf("list token error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	log.C(c).Infof("list token end")
}

func (tc *TokenController) Get(c *gin.Context) {
	log.C(c).Infof("get token begin")
	tokenId := c.Param(common.ParamUriTokenId)
	if tokenId == "" {
		log.C(c).Errorf("token id is empty")
		core.WriteResponse(c, errors.New("token id is required"), nil)
		return
	}

	resp, err := tc.b.Token().Get(c, tokenId)
	if err != nil {
		log.C(c).Errorf("get token error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	log.C(c).Infof("get token end")
}
