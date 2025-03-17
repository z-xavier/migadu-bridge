package bridges

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
)

func (bc *BridgesController) AddyAliases(c *gin.Context) (any, error) {
	log.C(c).Infof("AddyAliases begin")

	// https://app.addy.io/docs/#aliases-POSTapi-v1-aliases
	// 绑定 addy
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

	return nil, nil
}
