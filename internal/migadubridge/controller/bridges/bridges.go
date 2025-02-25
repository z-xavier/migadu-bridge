package bridges

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/core"
	"migadu-bridge/internal/pkg/log"
)

// BridgesController 定义了 controller 层需要实现的方法.
type BridgesController struct {
	b biz.IBiz
}

func New(ds store.IStore) *BridgesController {
	return &BridgesController{
		b: biz.NewBiz(ds),
	}
}

func (bc *BridgesController) AddyAliases(c *gin.Context) {
	log.C(c).Infof("AddyAliases begin")

	//var r v1.ListAliasReq
	//if err := c.ShouldBind(&r); err != nil {
	//	log.C(c).Errorf("AddyAliases request parse error: %s", err.Error())
	//	core.WriteResponse(c, errmsg.ErrBind.WithCause(err), nil)
	//	return
	//}
	//
	//resp, err := ac.b.Alias().List(c, &r)
	//if err != nil {
	//	log.C(c).Errorf("AddyAliases error: %s", err.Error())
	//	core.WriteResponse(c, err, nil)
	//	return
	//}

	core.WriteResponse(c, nil, nil)
	log.C(c).Infof("AddyAliases end")
}

func (bc *BridgesController) Simplelogin(c *gin.Context) {
	log.C(c).Infof("Simplelogin begin")

	//var r v1.ListAliasReq
	//if err := c.ShouldBind(&r); err != nil {
	//	log.C(c).Errorf("Simplelogin request parse error: %s", err.Error())
	//	core.WriteResponse(c, errmsg.ErrBind.WithCause(err), nil)
	//	return
	//}
	//
	//resp, err := ac.b.Alias().List(c, &r)
	//if err != nil {
	//	log.C(c).Errorf("Simplelogin error: %s", err.Error())
	//	core.WriteResponse(c, err, nil)
	//	return
	//}

	core.WriteResponse(c, nil, nil)
	log.C(c).Infof("Simplelogin end")
}
