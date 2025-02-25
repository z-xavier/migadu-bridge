package Aliaseses

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/core"
	"migadu-bridge/internal/pkg/log"
)

// AliasesController 定义了 controller 层需要实现的方法.
type AliasesController struct {
	b biz.IBiz
}

func New(ds store.IStore) *AliasesController {
	return &AliasesController{
		b: biz.NewBiz(ds),
	}
}

func (ac *AliasesController) List(c *gin.Context) {
	log.C(c).Infof("list aliasese begin")

	//var r v1.ListAliasesReq
	//if err := c.ShouldBind(&r); err != nil {
	//	log.C(c).Errorf("list call logs request parse error: %s", err.Error())
	//	core.WriteResponse(c, errmsg.ErrBind.WithCause(err), nil)
	//	return
	//}

	resp, err := ac.b.Alias().List(c)
	if err != nil {
		log.C(c).Errorf("list aliasese error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	log.C(c).Infof("list aliasese end")
}
