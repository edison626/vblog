package exception

import "fmt"

// New(5000, "令牌过期...")
func New(code int, format string, a ...any) *ApiException {
	return &ApiException{
		Code:    code,
		Message: fmt.Sprintf(format, a...),
	}
}

// 业务自定义异常
type ApiException struct {
	Code    int    `json:"code"`    //报错代码 如 404
	Message string `json:"message"` //报错信息
}

// 实现Error接口
func (e *ApiException) Error() string {
	return e.Message
}
