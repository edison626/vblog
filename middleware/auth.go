package middleware

import (
	"fmt"
	"net/http"

	"github.com/edison626/vblog/apps/token"
	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/exception"
	"github.com/edison626/vblog/ioc"
	"github.com/edison626/vblog/response"
	"github.com/gin-gonic/gin"
)

func NewTokenAuther() *TokenAuther {
	return &TokenAuther{
		//tk: ioc.GetController().Get(token.AppName).(token.Service),
		tk: ioc.Controller().Get(token.AppName).(token.Service),
	}
}

// 用来做鉴权的中间件
// 用于token 做鉴权的中间件
type TokenAuther struct {
	tk token.Service
	//给一个角色
	role user.Role
}

// 怎么鉴权？
// Gin中间件 - 需要实现 func（*Context）这函数
func (a *TokenAuther) Auth(c *gin.Context) {
	// 1. 获取Token
	at, err := c.Cookie(token.TOKEN_COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			response.Failed(c, token.CookieNotFound)
			return
		}
		response.Failed(c, err)
		return
	}

	// 2.调用Token模块来认证
	in := token.NewValiateToken(at)
	tk, err := a.tk.ValiateToken(c.Request.Context(), in)
	if err != nil {
		response.Failed(c, err)
		return
	}

	// 把鉴权后的 结果: tk, 放到请求的上下文, 方便后面的业务逻辑使用
	if c.Keys == nil {
		c.Keys = map[string]any{}
	}
	c.Keys[token.TOKEN_GIN_KEY_NAME] = tk
}

// 权限鉴定, 鉴权是在用户已经认证的情况之下进行的
// 判断当前用户的角色
func (a *TokenAuther) Perm(c *gin.Context) {
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

	fmt.Printf("user %s role %d \n", tk.UserName, tk.Role)

	// 如果是Admin则直接放行
	if tk.Role == user.ROLE_ADMIN {
		return
	}

	if tk.Role != a.role {
		response.Failed(c, exception.NewPermissionDeny("role %d not allow", tk.Role))
		return
	}
}

// 写带参数的 Gin中间件
func Required(r user.Role) gin.HandlerFunc {
	a := NewTokenAuther()
	a.role = r
	return a.Perm
}
