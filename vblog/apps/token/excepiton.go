package token

// 用于定义这个模块的独有异常

import "gitee.com/go-course/go12/vblog/exception"

var AuthFailed = exception.NewAuthFailed("用户名或者密码不正确")
