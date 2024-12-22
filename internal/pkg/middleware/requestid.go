package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"migadu-bridge/internal/pkg/common"
)

// RequestId 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 检查请求头中是否有 `X-Request-ID`，如果有则复用，没有则新建
		requestID := c.Request.Header.Get(common.XRequestIDKey)

		if requestID == "" {
			requestID = xid.New().String()
		}

		// 将 RequestID 保存在 gin.Context 中，方便后边程序使用
		c.Set(common.XRequestIDKey, requestID)

		// 将 RequestID 保存在 HTTP 返回头中，Header 的键为 `X-Request-ID`
		c.Writer.Header().Set(common.XRequestIDKey, requestID)
		c.Next()
	}
}
