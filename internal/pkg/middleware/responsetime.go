package middleware

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/common"
)

// ResponseTime 是一个 Gin 中间件，用来在每一个 HTTP 返回头中添加响应时间, Header 的键为 `X-Response-Time`
func ResponseTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		c.Writer.Header().Set(common.XResponseTime, strconv.FormatInt(time.Now().Unix(), 10))
	}
}
