package api

import (
	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/apps/token"
	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/common"
	"github.com/edison626/vblog/exception"
	"github.com/edison626/vblog/middleware"
	"github.com/edison626/vblog/response"
	"github.com/gin-gonic/gin"
)

// 需要把HandleFunc 添加到Root路由，定义 API ---> HandleFunc
// 可以选择把这个Handler上的HandleFunc都注册到路由上面
func (h *apiHandler) Registry(r gin.IRouter) {
	// r 是Gin的路由器
	v1 := r.Group("v1").GET("blogs")
	// 需要公开给访客访问 可以不用鉴权
	// GET /vblog/api/v1/blogs/
	v1.GET("/", h.QueryBlog)
	// GET /vblog/api/v1/blogs/43
	// url路径里面有变量: :id, 什么一个路径变量, r.Params("id") 43
	// /:id ---? /43   id = 43
	v1.GET("/:id", h.DescribeBlog)

	// 后台管理接口 需要认证
	v1.Use(middleware.NewTokenAuther().Auth)

	// POST /vblog/api/v1/blogs/
	// POST /vblog/api/v1/blogs/ -- HanlerFunc Chain: append(CreateBlog, middlewares.NewTokenAuther().Auth, ...)
	v1.POST("/", middleware.Required(user.ROLE_AUTHOR), h.CreateBlog)
	// PUT /vblog/api/v1/blogs/43
	v1.PUT("/:id", middleware.Required(user.ROLE_AUTHOR), h.UpdateBlog)
	// PATCH /vblog/api/v1/blogs/43
	v1.PATCH("/:id", middleware.Required(user.ROLE_AUTHOR), h.PatchBlog)
	// DELETE /vblog/api/v1/blogs/43
	v1.DELETE("/:id", middleware.Required(user.ROLE_AUTHOR), h.DeleteBlog)

	v1.POST("/:id/audit", middleware.Required(user.ROLE_AUDITOR), h.AuditBlog)
}

// 博客创建
func (h *apiHandler) CreateBlog(c *gin.Context) {
	// 怎么鉴权?
	// Gin中间件 func(*Context)
	// 	// 1. 获取Token
	// 	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	// 	if err != nil {
	// 		if err == http.ErrNoCookie {
	// 			response.Failed(c, token.CookieNotFound)
	// 			return
	// 		}
	// 		response.Failed(c, err)
	// 		return
	// 	}

	// 	// 2.调用Token模块来认证
	// 	in := token.NewValiateToken(at)
	// 	tk, err := a.tk.ValiateToken(c.Request.Context(), in)
	// 	if err != nil {
	// 		response.Failed(c, err)
	// 		return
	// 	}

	// 	// 把鉴权后的 结果: tk, 放到请求的上下文, 方便后面的业务逻辑使用
	// 	if c.Keys == nil {
	// 		c.Keys = map[string]any{}
	// 	}
	// 	c.Keys[token.TOKEN_GIN_KEY_NAME] = tk

	// 从Gin请求上下文中: c.Keys, 获取认证过后的鉴权结果
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	tk := tkObj.(*token.Token)

	in := blog.NewCreateBlogRequest()
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 从上下文中补充 用户信息
	in.CreateBy = tk.UserName
	ins, err := h.svc.CreateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 列表查询
// GET /vblog/api/v1/blogs/?page_size=1&page_number=20
func (h *apiHandler) QueryBlog(c *gin.Context) {

	// 从GIN 请求上下文中： c.keys，获取哦去认证的鉴全结果
	// tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	// fmt.Println(tkObj.(*token.Token).UserId)

	in := blog.NewQueryBlogRequest()
	in.ParsePageSize(c.Query("page_size"))
	in.ParsePageSize(c.Query("page_number"))
	switch c.Query("status") {
	case "draft":
		in.SetStatus(blog.STATUS_DRAFT)
	case "published":
		in.SetStatus(blog.STATUS_PUBLISHED)
	}

	set, err := h.svc.QueryBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, set)
}

// 详情查询
// GET /vblog/api/v1/blogs/43
func (h *apiHandler) DescribeBlog(c *gin.Context) {
	in := blog.NewDescribeBlogRequest(c.Param("id"))
	ins, err := h.svc.DescribeBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// 全量更新
// PUT /vblog/api/v1/blogs/43
// BODY  {} - 通过json 传回来的body参数
func (h *apiHandler) UpdateBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}
	// 优化这部分逻辑
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	if tkObj == nil {
		response.Failed(c, exception.NewPermissionDeny("token not found"))
		return
	}

	tk, ok := tkObj.(*token.Token)
	if !ok {
		response.Failed(c, exception.NewPermissionDeny("token not an *token.Token"))
		return
	}

	in.Scope = &common.Scope{
		Username: tk.UserName,
	}
	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, ins)
}

// 增量更新
// PATCH /vblog/api/v1/blogs/43
func (h *apiHandler) PatchBlog(c *gin.Context) {
	in := blog.NewPutUpdateBlogRequest(c.Param("id"))
	err := c.BindJSON(in.CreateBlogRequest)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 优化这部分逻辑
	tkObj := c.Keys[token.TOKEN_GIN_KEY_NAME]
	if tkObj == nil {
		response.Failed(c, exception.NewPermissionDeny("token not found"))
		return
	}

	tk, ok := tkObj.(*token.Token)
	if !ok {
		response.Failed(c, exception.NewPermissionDeny("token not an *token.Token"))
		return
	}

	in.Scope = &common.Scope{
		Username: tk.UserName,
	}

	ins, err := h.svc.UpdateBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}

// 删除
// DELETE /vblog/api/v1/blogs/43
func (h *apiHandler) DeleteBlog(c *gin.Context) {
	in := blog.NewDeleteBlogRequest(c.Param("id"))

	err := h.svc.DeleteBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}
	response.Success(c, "ok")
}

// 审核该文章 是否合法
// POST /vblog/api/v1/blogs/43/audit
func (h *apiHandler) AuditBlog(c *gin.Context) {
	in := blog.NewAuditBlogRequest(c.Param("id"))
	err := c.BindJSON(in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	ins, err := h.svc.AuditBlog(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	response.Success(c, ins)
}
