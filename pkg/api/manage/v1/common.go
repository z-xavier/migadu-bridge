package v1

// Response 统一返回消息.
type Response struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`

	// Data 包含了
	Data any `json:"data"`
}
