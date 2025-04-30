// Package aliases implements the aliases API endpoints
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

// List godoc
// @Summary List aliases
// @Description Get a list of aliases with pagination
// @Tags aliases
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1) minimum(1)
// @Param pageSize query int false "Page size" default(10) minimum(1) maximum(200)
// @Param orderBy query []string false "Order by fields"
// @Success 200 {object} v1.ListAliasResp "Success"
// @Failure 400 {object} v1.Response "Bad request"
// @Failure 500 {object} v1.Response "Internal server error"
// @Router /api/v1/aliases [get]
func (ac *Controller) List(c *gin.Context) (any, error) {
	log.C(c).Info("list aliasese begin")

	var r v1.ListAliasReq
	if err := c.ShouldBind(&r); err != nil {
		log.C(c).WithError(err).Error("list aliasese request parse")
		return nil, errmsg.ErrBind.WithCause(err)
	}

	return ac.b.Alias().List(c, &r)
}
