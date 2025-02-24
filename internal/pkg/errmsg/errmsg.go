package errmsg

import (
	"errors"
	"fmt"
)

// Errmsg 定义了 migadu-provider 使用的错误类型.
type Errmsg struct {
	HTTP    int
	Code    string
	Message string
}

// Error 实现 error 接口中的 `Error` 方法.
func (err *Errmsg) Error() string {
	return err.Message
}

// SetMessage 设置 Errno 类型错误中的 Message 字段.
func (err *Errmsg) SetMessage(format string, args ...interface{}) *Errmsg {
	err.Message = fmt.Sprintf(format, args...)
	return err
}

// WithCause 设置 Errno 类型错误中的 Message 字段.
func (err *Errmsg) WithCause(cause error) *Errmsg {
	err.Message = cause.Error()
	return err
}

// Decode 尝试从 err 中解析出业务错误码和错误信息.
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	var typed *Errmsg
	switch {
	case errors.As(err, &typed):
		return typed.HTTP, typed.Code, typed.Message
	default:
		// 默认返回未知错误码和错误信息. 该错误代表服务端出错
		return InternalServerError.HTTP, InternalServerError.Code, err.Error()
	}
}
