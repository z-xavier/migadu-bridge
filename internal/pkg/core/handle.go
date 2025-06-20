package core

import (
	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/log"
)

type handler func(c *gin.Context) (any, error)

func HandleResult(h handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp, err := h(c)
		if err != nil {
			log.C(c).WithError(err).Error("handle result error")
			WriteResponse(c, err, nil)
			return
		}
		WriteResponse(c, nil, resp)
	}
}
