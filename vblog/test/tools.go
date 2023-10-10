package test

import (
	"gitee.com/go-course/go12/vblog/conf"
	"gitee.com/go-course/go12/vblog/ioc"

	// 注册对象
	_ "gitee.com/go-course/go12/vblog/apps"
)

// 设置单元测试的配置和环境
func DevelopmentSetup() {
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 对象的初始化
	if err := ioc.Controller().Init(); err != nil {
		panic(err)
	}
}
