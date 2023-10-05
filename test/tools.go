package test

import (
	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/ioc"
	//注册对象
	_ "github.com/edison626/vblog/apps"
)

// 设置单元测试的配置和环境
func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	//对象的初始化
	if err := ioc.Controller().Init(); err!= nil{
		panic(err)
	}
	
}
