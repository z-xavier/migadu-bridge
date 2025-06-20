package core

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"migadu-bridge/internal/pkg/errmsg"
	v1 "migadu-bridge/pkg/api/manage/v1"
)

// WriteResponse 将错误或响应数据写入 HTTP 响应主体。
// WriteResponse 使用 errno.Decode 方法，根据错误类型，尝试从 err 中提取业务错误码和错误信息.
func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		hcode, code, message := errmsg.Decode(err)
		c.JSON(hcode, v1.Response{
			Code:    code,
			Message: message,
			Data:    make(map[string]any),
		})

		return
	}

	c.JSON(http.StatusOK, v1.Response{
		Data: data,
	})
}
