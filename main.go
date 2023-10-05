package main

import (
	"fmt"
	"os"

	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/ioc"
	"github.com/gin-gonic/gin"

	//注册对象
	_ "github.com/edison626/vblog/apps"

	tokenApiHandler "github.com/edison626/vblog/apps/token/api"
)

func main() {
	//1. 加载配置
	err := conf.LoadConfigFromToml("etc/application.toml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//2. 初始化控制
	// 2.1 user controller
	//userServiceImpl := userImpl.NewUserServiceImpl()

	// 2.2 token controller
	//tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)

	//通过Ioc 来完成依赖装载，完成了依赖的倒置（ioc依赖对象注册）
	if err := ioc.Controller().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 2.3 token api handler
	tkApiHandler := tokenApiHandler.NewTokenApiHandler()

	//3. 启动http协议服务器, 注册 handler路由
	r := gin.Default()
	tkApiHandler.Registry(r.Group("/api/vblog"))

	// 启动协议服务器
	addr := conf.C().App.HttpAddr() //在config 配置 app
	fmt.Printf("HTTP API监听地址: %s", addr)
	err = r.Run(addr)
	fmt.Println(err)
}
