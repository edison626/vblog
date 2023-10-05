package impl_test

import (
	"context"
	"testing"

	"github.com/edison626/vblog/apps/token"
	"github.com/edison626/vblog/ioc"
	"github.com/edison626/vblog/test"
)

var (
	//要测试的是token service
	tokenSvc token.Service
	ctx      = context.Background()
)

// 测试登陆后，会颁发token存到数据库
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

// 测试token 对应是否正确
func TestValiateToken(t *testing.T) {
	//复制mysql access toke 做测试
	req := token.NewValiateToken("ck957fof79qj20cirf80")
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
