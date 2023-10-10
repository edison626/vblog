package impl_test

import (
	"context"
	"testing"

	"gitee.com/go-course/go12/vblog/apps/token"
	"gitee.com/go-course/go12/vblog/ioc"
	"gitee.com/go-course/go12/vblog/test"
)

var (
	tokenSvc token.Service
	ctx      = context.Background()
)

func TestLogin(t *testing.T) {
	req := token.NewLoginRequest()
	req.Username = "admin"
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
	tokenSvc = ioc.Controller().Get(token.AppName).(token.Service)
}
