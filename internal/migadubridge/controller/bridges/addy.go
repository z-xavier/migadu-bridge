package bridges

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/pkg/api/manage/bridge/addy"
)

// AddyAliases godoc
// @Summary Create an alias for Addy.io
// @Description Create a new email alias using the Addy.io API
// @Tags bridges
// @Accept json
// @Produce json
// @Param Authorization header string true "API Token"
// @Param request body addy.CreateAliasReq true "Alias creation request"
// @Success 201 {object} addy.CreateAliasResp "Created"
// @Failure 400 {object} addy.ErrorResp "Bad request"
// @Failure 401 {object} addy.ErrorResp "Unauthorized"
// @Failure 500 {object} addy.ErrorResp "Internal server error"
// @Router /api/v1/aliases [post]
func (bc *Controller) AddyAliases(c *gin.Context) {
	log.C(c).Info("AddyAliases begin")

	// https://app.addy.io/docs/#aliases-POSTapi-v1-aliases
	var r addy.CreateAliasReq
	if err := c.ShouldBindHeader(&r); err != nil {
		log.C(c).WithError(err).Error("AddyAliases request parse")
		c.JSON(http.StatusUnauthorized, addy.ErrorResp{Error: err.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		log.C(c).WithError(err).Error("AddyAliases request parse")
		c.JSON(http.StatusBadRequest, addy.ErrorResp{Error: err.Error()})
		return
	}

	resp, errCode, err := bc.b.Bridge().AddyAliases(c, &r)
	if err != nil {
		log.C(c).WithError(err).Error("AddyAliases")
		c.JSON(errCode, addy.ErrorResp{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
	return
}
