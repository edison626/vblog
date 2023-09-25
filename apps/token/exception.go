package token

// 用于定义这个模块的独有异常

import "github.com/edison626/vblog/exception"

// 被impl 调用返回信息
var AuthFailed = exception.NewAuthFailed("用户名或者密码不正确")
