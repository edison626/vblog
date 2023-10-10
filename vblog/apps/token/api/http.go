package api

import (
	"gitee.com/go-course/go12/vblog/apps/token"
	"gitee.com/go-course/go12/vblog/ioc"
	"gitee.com/go-course/go12/vblog/response"
	"github.com/gin-gonic/gin"
)

func init() {
	ioc.ApiHandler().Registry(&TokenApiHandler{})
}

// 不适用接口, 直接定义Gin的一个handlers
// 什么是Gin的Handler  HandlerFunc
// HandlerFunc defines the handler used by gin middleware as return value.
// type HandlerFunc func(*Context)
// HandleFunc 只是定义 如何处理 HTTP 的请求与响应

type TokenApiHandler struct {
	// 依赖控制器
	svc token.Service
}

func (t *TokenApiHandler) Name() string {
	return token.AppName
}

func (t *TokenApiHandler) Init() error {
	t.svc = ioc.Controller().Get(token.AppName).(token.Service)
	return nil
}

// 需要把HandleFunc 添加到Root路由，定义 API ---> HandleFunc
// 可以选择把这个Handler上的HandleFunc都注册到路由上面
func (h *TokenApiHandler) Registry(r gin.IRouter) {
	// r 是Gin的路由器
	v1 := r.Group("v1")
	v1.POST("/tokens/", h.Login)
	v1.DELETE("/tokens/", h.Logout)
}

// Login HandleFunc
func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1. 获取用户的请求参数， 参数在Body里面
	// 一定要使用JSON
	req := token.NewLoginRequest()

	// json.unmarsal
	// http boyd ---> LoginRequest Object
	err := c.BindJSON(req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 2. 执行逻辑
	// 把http 协议的请求 ---> 控制器的请求
	ins, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 3. 返回响应
	response.Success(c, ins)
}

// Logout HandleFunc
func (h *TokenApiHandler) Logout(*gin.Context) {

}
