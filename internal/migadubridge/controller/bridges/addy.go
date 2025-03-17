package bridges

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
	"migadu-bridge/pkg/api/manage/bridge/addy"
)

func (bc *BridgesController) AddyAliases(c *gin.Context) {
	log.C(c).Infof("AddyAliases begin")

	// https://app.addy.io/docs/#aliases-POSTapi-v1-aliases
	var r addy.CreateAliasReq
	if err := c.ShouldBindHeader(&r); err != nil {
		log.C(c).Errorf("AddyAliases request parse error: %s", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindBodyWithJSON(&r); err != nil {
		log.C(c).Errorf("AddyAliases request parse error: %s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, errCode, err := bc.b.Bridge().AddyAliases(c, &r)
	if err != nil {
		log.C(c).Errorf("AddyAliases error: %s", err.Error())
		c.JSON(errCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
	return
}
