package ioc

// 专门用于注册Controller 对象
func Controller() *IocContainter {
	return controllerContainer
}

// ioc 注册表对象，全局只有
var controllerContainer = &IocContainter{
	store: map[string]IocObject{},
}

// 定义逻辑 - 一个对象的注册表
type IocContainter struct {
	//采用Map来保持对象注册
	store map[string]IocObject
}

// 负责初始化所有的对象
func (c *IocContainter) Init() error {
	for _, obj := range c.store {
		if err := obj.Init(); err != nil {
			return err
		}
	}
	return nil
}

// 对象池
// 把对象注册
func (c *IocContainter) Registry(obj IocObject) {
	c.store[obj.Name()] = obj
}

// 获取
func (c *IocContainter) Get(name string) any {
	return c.store[name]
}
