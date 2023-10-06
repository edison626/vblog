package ioc

// 专门用于注册Controller对象
func ApiHandler() *IocContainter {
	return apiHandlerContainer
}

// ioc 注册表对象, 全局只有
var apiHandlerContainer = &IocContainter{
	store: map[string]IocObject{},
}
