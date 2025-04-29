package middleware

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"migadu-bridge/internal/pkg/common"
)

// RequestId 是一个 Gin 中间件，用来在每一个 HTTP 请求的 context, response 中注入 `X-Request-ID` 键值对.
func RequestId() gin.HandlerFunc {
	return requestid.New(
		requestid.WithGenerator(func() string {
			return xid.New().String()
		}),
		requestid.WithCustomHeaderStrKey(common.XRequestIDKey),
		requestid.WithHandler(func(c *gin.Context, requestID string) {
			c.Set(common.XRequestIDKey, requestID)
		}),
	)
}
