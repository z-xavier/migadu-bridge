package tokens

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
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
	if err := c.ShouldBindJSON(&r); err != nil {
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

func (tc *TokenController) Delete(c *gin.Context) {}

func (tc *TokenController) Update(c *gin.Context) {}

func (tc *TokenController) Patch(c *gin.Context) {}

func (tc *TokenController) List(ctx *gin.Context) {}

func (tc *TokenController) Get(ctx *gin.Context) {}
