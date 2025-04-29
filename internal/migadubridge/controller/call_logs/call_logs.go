// Package call_logs implements the call logs API endpoints
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

// List godoc
// @Summary List call logs
// @Description Get a list of call logs with pagination and filtering
// @Tags call-logs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param pageSize query int false "Page size" default(10) minimum(1) maximum(200)
// @Param orderBy query []string false "Order by fields"
// @Param targetEmail query string false "Filter by target email"
// @Param mockProvider query string false "Filter by mock provider"
// @Param requestPath query string false "Filter by request path"
// @Param requestIp query string false "Filter by request IP"
// @Param requestAtBegin query int64 false "Filter by request time (begin)"
// @Param requestAtEnd query int64 false "Filter by request time (end)"
// @Success 200 {object} v1.ListCallLogResp "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/calllogs [get]
func (cc *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Info("list call logs begin")

	var r v1.ListCallLogReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("list call logs request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return cc.b.CallLog().List(c, &r)
}
