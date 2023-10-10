package impl_test

import (
	"context"
	"testing"

	"gitee.com/go-course/go12/vblog/apps/user"
	"gitee.com/go-course/go12/vblog/exception"
	"gitee.com/go-course/go12/vblog/ioc"
	"gitee.com/go-course/go12/vblog/test"
)

var (
	userSvc user.Service
	ctx     = context.Background()
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin"
	req.Password = "123456"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: 9,
	})
	if err != nil {
		if v, ok := err.(*exception.ApiException); ok {
			t.Fatal("error code", v.BizCode)
		}
		t.Fatal(err)
	}
}

func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("9")
	ins, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
func TestDescribeUserRequestByName(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin")
	ins, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)

	err = ins.CheckPassword("123456")
	if err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
	// 取消对象: ioc.Controller().Get(user.AppName)
	// 断言为接口来使用(只使用对象接口提供出来的能力)
	userSvc = ioc.Controller().Get(user.AppName).(user.Service)
}
