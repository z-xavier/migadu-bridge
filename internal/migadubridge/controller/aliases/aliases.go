package aliases

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

// Controller 定义了 controller 层需要实现的方法.
type Controller struct {
	b biz.IBiz
}

func New(ds store.IStore) *Controller {
	return &Controller{
		b: biz.NewBiz(ds),
	}
}

func (ac *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Infof("list aliasese begin")

	var r v1.ListAliasReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("list aliasese request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return ac.b.Alias().List(c, &r)
}
