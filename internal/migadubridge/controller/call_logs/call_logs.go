package call_logs

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

func (cc *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Infof("list call logs begin")

	var r v1.ListCallLogReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("list call logs request parse error: %s", err.Error())
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return cc.b.CallLog().List(c, &r)
}
