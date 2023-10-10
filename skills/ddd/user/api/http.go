package api

import (
	"gitee.com/go-course/go12/skills/ddd/user"
	"github.com/gin-gonic/gin"
)

func NewHandler(svc user.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// 定义一个对象, 就负责实现接口暴露(Restful)
type Handler struct {
	// 依赖一个svc实现来
	svc user.Service
}

// 实现 CreateUser接口
func (h *Handler) CreateUser(c *gin.Context) {
	h.svc.CreateUser(c.Request.Context(), &user.CreateUserRequest{})
}
