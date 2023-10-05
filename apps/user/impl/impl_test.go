package impl_test

import (
	"context"
	"testing"

	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/exception"
	"github.com/edison626/vblog/ioc"
	"github.com/edison626/vblog/test"
)

// 测试环境 - 全局变量
var (
	userSvc user.Service
	ctx     = context.Background()
)

func TestCreateUser(t *testing.T) {
	req := user.NewCreateUserRequest()
	req.Username = "admin2"
	req.Password = "123456"
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestDeleteUser(t *testing.T) {
	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
		Id: 18, //删除ID8
	})
	//自定义包装，调用 exception
	if err != nil {
		//使用断言的方式
		if v, ok := err.(*exception.ApiException); ok {
			t.Fatal("自定义 - error code", v.BizCode)
		}
		t.Fatal(err)
	}
}

// 测试查询 - name 查询
// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
func TestDescribeUserRequestByName(t *testing.T) {
	req := user.NewDescribeUserRequestByUsername("admin1")
	ins, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// 测试查询 - ID 查询
// 修改ID 和 mysql 的id 一样
func TestDescribeUserRequestById(t *testing.T) {
	req := user.NewDescribeUserRequestById("19")
	ins, err := userSvc.DescribeUserRequest(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// userSvc 被初始化为 impl.UserServiceImpl 的一个新实例的指针。之后，在其他函数（比如测试函数）中就可以使用这个已经初始化的 userSvc 变量了。
func init() {
	test.DevelopmentSetup()

	//取消的对象 ioc.Controller().Get(user.AppName)
	//断言为接口来使用（只使用对象接口提供出来的能力）
	userSvc = ioc.Controller().Get(user.AppName).(user.Service)

}
