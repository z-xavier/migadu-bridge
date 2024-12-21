package errmsg

var (
	// OK 代表请求成功.
	OK = &Errmsg{HTTP: 200, Code: "", Message: ""}

	// InternalServerError 表示所有未知的服务器端错误.
	InternalServerError = &Errmsg{HTTP: 500, Code: "InternalError", Message: "Internal server error."}

	// ErrPageNotFound 表示路由不匹配错误.
	ErrPageNotFound = &Errmsg{HTTP: 404, Code: "ResourceNotFound.PageNotFound", Message: "Page not found."}

	// ErrBind 表示参数绑定错误.
	ErrBind = &Errmsg{HTTP: 400, Code: "InvalidParameter.BindError", Message: "Error occurred while binding the request body to the struct."}

	// ErrInvalidParameter 表示所有验证失败的错误.
	ErrInvalidParameter = &Errmsg{HTTP: 400, Code: "InvalidParameter", Message: "Parameter verification failed."}

	// ErrSignToken 表示签发 JWT Token 时出错.
	ErrSignToken = &Errmsg{HTTP: 401, Code: "AuthFailure.SignTokenError", Message: "Error occurred while signing the JSON web token."}

	// ErrTokenInvalid 表示 JWT Token 格式错误.
	ErrTokenInvalid = &Errmsg{HTTP: 401, Code: "AuthFailure.TokenInvalid", Message: "Token was invalid."}

	// ErrUnauthorized 表示请求没有被授权.
	ErrUnauthorized = &Errmsg{HTTP: 401, Code: "AuthFailure.Unauthorized", Message: "Unauthorized."}
)
