package handler

import (
	"gitee.com/go-course/go12/skills/mvc/controller"
	"gitee.com/go-course/go12/skills/mvc/model"
	"github.com/gin-gonic/gin"
)

// 实现 CreateUser接口
func CreateUser(c *gin.Context) {
	// http 请求
	// c.Request

	// 从http 协议读取用户请求的参数
	req := &model.User{}
	c.Bind(req)

	// 调用Controller 来执行业务逻辑
	controller.CreateUser(req)

	// http 响应
	// c.Writer
	c.JSON(0, req)
}
