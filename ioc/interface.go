package ioc

//定义注册进来的对象的约束条件

type IocObject interface {
	// 对象的初始化
	Init() error
	//对象名称
	Name() string
}
