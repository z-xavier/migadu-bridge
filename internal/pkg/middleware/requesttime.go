package middleware

import (
	"time"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/common"
)

// RequestTime 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context 中注入 `X-Request-Time` 键值对.
func RequestTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(common.XRequestTime, time.Now())
	}
}
