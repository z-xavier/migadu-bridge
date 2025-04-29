package bridges

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

// SLAliasRandomNew godoc
// @Summary Create a random alias for SimpleLogin
// @Description Create a new random email alias using the SimpleLogin API
// @Tags bridges
// @Accept json
// @Produce json
// @Param Authentication header string true "API Key"
// @Param mode query string false "Alias mode"
// @Param request body sl.AliasRandomNewReq true "Alias creation request"
// @Success 201 {object} sl.Alias "Created"
// @Failure 400 {object} sl.ErrorResp "Bad request"
// @Failure 401 {object} sl.ErrorResp "Unauthorized"
// @Failure 500 {object} sl.ErrorResp "Internal server error"
// @Router /api/alias/random/new [post]
func (bc *Controller) SLAliasRandomNew(c *gin.Context) {
	log.C(c).Info("SLAliasRandomNew begin")

	// https://github.com/simple-login/app/blob/master/docs/api.md#post-apialiasrandomnew
	var r sl.AliasRandomNewReq
	if err := c.ShouldBindHeader(&r); err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew request parse")
		c.JSON(http.StatusUnauthorized, sl.ErrorResp{Error: err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew request parse")
		c.JSON(http.StatusBadRequest, sl.ErrorResp{Error: err.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew request parse")
		c.JSON(http.StatusBadRequest, sl.ErrorResp{Error: err.Error()})
		return
	}

	resp, errCode, err := bc.b.Bridge().SLAliasRandomNew(c, &r)
	if err != nil {
		log.C(c).WithError(err).Error("SLAliasRandomNew")
		c.JSON(errCode, sl.ErrorResp{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
	return
}
