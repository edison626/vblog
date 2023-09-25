package impl_test

import (
	"context"
	"testing"

	"github.com/edison626/vblog/apps/token"
	"github.com/edison626/vblog/apps/token/impl"
	userImpl "github.com/edison626/vblog/apps/user/impl"
	"github.com/edison626/vblog/test"
)

var (
	tokenSvc *impl.TokenServiceImpl
	ctx      = context.Background()
)

func TestLogin(t *testing.T) {
	req := token.NewLoginRequest()
	req.Username = "admin2"
	req.Password = "123456"
	tk, err := tokenSvc.Login(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func TestValiateToken(t *testing.T) {
	req := token.NewValiateToken("ck2m1hlmjd0jbb0g3ip0")
	tk, err := tokenSvc.ValiateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tk)
}

func init() {
	test.DevelopmentSetup()

	// 依赖另一个实现类
	tokenSvc = impl.NewTokenServiceImpl(userImpl.NewUserServiceImpl())
}
