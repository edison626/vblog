package ioc

// 专门用于注册Controller对象
func Controller() *IocContainter {
	return controllerContainer
}

// ioc 注册表对象, 全局只有
var controllerContainer = &IocContainter{
	store: map[string]IocObject{},
}
