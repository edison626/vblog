package main

import (
	"fmt"
	"os"

	"gitee.com/go-course/go12/vblog/conf"
	"gitee.com/go-course/go12/vblog/ioc"
	"github.com/gin-gonic/gin"

	// 注册对象
	_ "gitee.com/go-course/go12/vblog/apps"
)

func main() {
	//1. 加载配置
	err := conf.LoadConfigFromToml("etc/application.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//2. 初始化控制
	// 收到传递依赖关系：收到管理对象依赖
	// 2.1 user controller
	// userServiceImpl := userImpl.NewUserServiceImpl()
	// 2.2 token controller
	// tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)
	// ....
	//  通过Ioc来完成依赖的装载, 完成了依赖的倒置(ioc 依赖对象注册)
	if err := ioc.Controller().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 初始化Api Handler
	if err := ioc.ApiHandler().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//3. 启动http协议服务器, 注册 handler路由
	r := gin.Default()
	ioc.ApiHandler().RouteRegistry(r.Group("/api/vblog"))

	// 启动协议服务器
	addr := conf.C().App.HttpAddr()
	fmt.Printf("HTTP API监听地址: %s", addr)
	err = r.Run(addr)
	fmt.Println(err)
}
