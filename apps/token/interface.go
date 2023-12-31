package token

import "context"

// ioc 为这模块定义个名
const (
	AppName = "token"
)

type Service interface {
	// 登录接口(颁发Token)
	Login(context.Context, *LoginRequest) (*Token, error)
	// 退出接口(销毁Token)
	Logout(context.Context, *LogoutRequest) error

	// 校验Token 是给内部中间层使用 身份校验层
	// 校验完后返回Token, 通过Token获取 用户信息
	ValiateToken(context.Context, *ValiateToken) (*Token, error)
}

// 登陆函数
func NewLoginRequest() *LoginRequest {
	return &LoginRequest{}
}

type LoginRequest struct {
	Username string
	Password string
}

// 万一的Token泄露, 不知道refresh_token，也没法推出
type LogoutRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewValiateToken(at string) *ValiateToken {
	return &ValiateToken{
		AccessToken: at,
	}
}

type ValiateToken struct {
	AccessToken string `json:"access_token"`
}
