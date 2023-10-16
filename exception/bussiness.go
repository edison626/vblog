package exception

// 业务自定义码 404 - 资源没找到
func NewNotFound(format string, a ...any) *ApiException {
	return New(404, format, a...)
}

func IsNotFound(err error) bool {
	if e, ok := err.(*ApiException); ok {
		if e.BizCode == 404 {
			return true
		}
	}

	return false
}

// 认证失败
func NewAuthFailed(format string, a ...any) *ApiException {
	return New(5000, format, a...)
}

func NewPermissionDeny(format string, a ...any) *ApiException {
	return New(5100, format, a...)
}

func NewTokenExpired(format string, a ...any) *ApiException {
	return New(5001, format, a...)
}
