package ioc

import "github.com/gin-gonic/gin"

// 定义接口 就是定义逻辑
// 定义一个对象的注册表，IocContainter
type IocContainter struct {
	// 采用Map来保持对象注册
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

func (c *IocContainter) Registry(obj IocObject) {
	c.store[obj.Name()] = obj
}

func (c *IocContainter) Get(name string) any {
	return c.store[name]
}

// Gin
type GinApiHandler interface {
	Registry(r gin.IRouter)
}

// 管理者所有的对象(Api Handler)
// 把每个 ApiHandler的路由注册给Root Router
func (c *IocContainter) RouteRegistry(r gin.IRouter) {
	// 找到被托管的APIHandler
	for _, obj := range c.store {
		// 怎么来判断这个对象是一个APIHandler对象
		if api, ok := obj.(GinApiHandler); ok {
			api.Registry(r)
		}
	}
}
