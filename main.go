package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/ioc"
	"github.com/edison626/vblog/protocol"

	//注册对象
	_ "github.com/edison626/vblog/apps"
	//tokenApiHandler "github.com/edison626/vblog/apps/token/api"
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

	// 初始化Api Handler
	if err := ioc.ApiHandler().Init(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 跑在后台的http server
	httpServer := protocol.NewHttpServer()
	go func() {
		if err := httpServer.Run(); err != nil {
			fmt.Printf("start http server error, %s\n", err)
		}
	}()
	//3. 启动http协议服务器, 注册 handler路由
	// r := gin.Default()
	// ioc.ApiHandler().RouteRegistry(r.Group("/api/vblog"))

	// // 启动协议服务器
	// addr := conf.C().App.HttpAddr()
	// fmt.Printf("HTTP API监听地址: %s", addr)
	// err = r.Run(addr)

	// fmt.Println(err)
	// fmt.Println("清理工作")

	// 处理信号量
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

	// 等待信号的到来
	sn := <-ch
	fmt.Println(sn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Close(ctx)

	//
	fmt.Println("清理工作")
}
