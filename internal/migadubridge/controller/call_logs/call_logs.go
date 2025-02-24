package call_logs

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/migadubridge/biz"
	"migadu-bridge/internal/migadubridge/store"
	"migadu-bridge/internal/pkg/core"
	"migadu-bridge/internal/pkg/errmsg"
	"migadu-bridge/internal/pkg/log"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

// CallLogsController 定义了 controller 层需要实现的方法.
type CallLogsController struct {
	b biz.IBiz
}

func New(ds store.IStore) *CallLogsController {
	return &CallLogsController{
		b: biz.NewBiz(ds),
	}
}

func (cc *CallLogsController) List(c *gin.Context) {
	log.C(c).Infof("list call logs begin")

	var r v1.ListCallLogReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).Errorf("list call logs request parse error: %s", err.Error())
		core.WriteResponse(c, errmsg.ErrBind.WithCause(err), nil)
		return
	}

	resp, err := cc.b.CallLog().List(c, &r)
	if err != nil {
		log.C(c).Errorf("list call logs error: %s", err.Error())
		core.WriteResponse(c, err, nil)
		return
	}

	core.WriteResponse(c, nil, resp)
	log.C(c).Infof("list call logs end")
}
