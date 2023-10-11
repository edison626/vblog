package exception

import (
	"fmt"
	"net/http"
)

// New(5000, "令牌过期...")
func New(code int, format string, a ...any) *ApiException {
	return &ApiException{
		BizCode:  code,
		Message:  fmt.Sprintf(format, a...),
		HttpCode: http.StatusOK,
	}
}

// 业务自定义异常
type ApiException struct {
	//业务异常
	BizCode  int    `json:"code"`    //报错代码 如 404
	Message  string `json:"message"` //报错信息
	Data     any    `json:"data"`
	HttpCode int    `json:"http_code"`
}

// 实现Error接口
func (e *ApiException) Error() string {
	return e.Message
}
