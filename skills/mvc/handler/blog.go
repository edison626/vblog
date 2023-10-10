package handler

import (
	"gitee.com/go-course/go12/skills/mvc/model"
	"github.com/gin-gonic/gin"
)

// 实现 CreateBlog 接口
func CreateBlog(c *gin.Context) {
	// http 请求
	// c.Request

	// 从http 协议读取用户请求的参数
	req := &model.Blog{}
	c.Bind(req)

	// 调用Controller 来执行业务逻辑
	// controller.CreateBlog(req)

	// http 响应
	// c.Writer
	c.JSON(0, req)
}
