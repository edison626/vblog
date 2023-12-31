package impl

import (
	"context"
	"fmt"

	"github.com/edison626/vblog/apps/token"
	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/exception"
	"github.com/edison626/vblog/ioc"
	"gorm.io/gorm"
)

// 传递用户接口的实例
func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{
		//初始化db
		db: conf.C().MySQL.GetConn().Debug(),
		//传user 服务实现
		user: ioc.Controller().Get(user.AppName).(user.Service),
	}
}

func init() {
	ioc.Controller().Registry(&TokenServiceImpl{})
}

type TokenServiceImpl struct {
	// db
	db *gorm.DB
	// 依赖User模块, 直接操作user模块的数据库(users)?
	// 这里需要依赖另一个业务领域: 用户管理领域
	user user.Service
}

func (i *TokenServiceImpl) Init() error {
	//初始化db
	i.db = conf.C().MySQL.GetConn().Debug()
	//传user 服务实现
	//拿到的对象，在main 进行初始化好了
	i.user = ioc.Controller().Get(user.AppName).(user.Service)
	return nil
}

// token 模块控制器的实现
func (i *TokenServiceImpl) Name() string {
	return token.AppName
}

// 登录接口(颁发Token)
func (i *TokenServiceImpl) Login(
	ctx context.Context, req *token.LoginRequest) (
	*token.Token, error) {
	// 1. 查询用户
	uReq := user.NewDescribeUserRequestByUsername(req.Username)
	u, err := i.user.DescribeUserRequest(ctx, uReq)
	if err != nil {
		if exception.IsNotFound(err) {
			// 调用 exception 里的AuthFail信息
			return nil, token.AuthFailed
		}
		return nil, err
	}

	// 2. 比对密码
	err = u.CheckPassword(req.Password)
	if err != nil {
		return nil, token.AuthFailed
	}

	// 3. 颁发token
	tk := token.NewToken()
	tk.UserId = u.Id
	tk.UserName = u.Username

	// 4. 保存Token
	if err := i.db.
		WithContext(ctx).
		Create(tk).
		Error; err != nil {
		return nil, err
	}

	// 避免同一个用户多次登录
	// 4. 颁发成功后  把之前的Token标记为失效,作业
	return tk, nil
}

// 校验Token 是给内部中间层使用 身份校验层
func (i *TokenServiceImpl) ValiateToken(
	ctx context.Context,
	req *token.ValiateToken) (*token.Token, error) {
	// 1. 查询Token (是不是我们这个系统颁发的)
	tk := token.NewToken()
	err := i.db.
		WithContext(ctx).
		Where("access_token = ?", req.AccessToken).
		First(tk).
		Error
	if err != nil {
		return nil, err
	}

	// 2. 判断Token的合法性:
	// 2.1 判断Ak是否过期
	if err := tk.IsExpired(); err != nil {
		return nil, err
	}

	// 补充用户信息, 只补充了用户的角色
	uDesc := user.NewDescribeUserRequestById(fmt.Sprintf("%d", tk.UserId))
	u, err := i.user.DescribeUserRequest(ctx, uDesc)
	if err != nil {
		return nil, err
	}
	tk.Role = u.Role

	return tk, nil
}

// 退出接口(销毁Token)
func (i *TokenServiceImpl) Logout(
	ctx context.Context,
	req *token.LogoutRequest) error {
	return nil
}
