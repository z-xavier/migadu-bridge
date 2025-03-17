package bridges

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/pkg/api/manage/bridge/sl"
)

func (bc *BridgesController) SLAliasRandomNew(c *gin.Context) {
	log.C(c).Infof("SLAliasRandomNew begin")

	// https://github.com/simple-login/app/blob/master/docs/api.md#post-apialiasrandomnew
	var r sl.AliasRandomNewReq
	if err := c.ShouldBindHeader(&r); err != nil {
		log.C(c).Errorf("SLAliasRandomNew request parse error: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindQuery(&r); err != nil {
		log.C(c).Errorf("SLAliasRandomNew request parse error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		log.C(c).Errorf("SLAliasRandomNew request parse error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, errCode, err := bc.b.Bridge().SLAliasRandomNew(c, &r)
	if err != nil {
		log.C(c).Errorf("SLAliasRandomNew error: %s", err.Error())
		c.JSON(errCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
	return
}
