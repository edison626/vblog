package impl_test

import (
	"context"
	"testing"

	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/apps/user/impl"
	"github.com/edison626/vblog/test"
)

// 测试环境 - 全局变量
var (
	userSvc *impl.UserServiceImpl
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

// func TestDeleteUser(t *testing.T) {
// 	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
// func TestDescribeUserRequestByName(t *testing.T) {
// 	req := user.NewDescribeUserRequestByUsername("admin")
// 	ins, err := userSvc.DescribeUserRequest(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(ins)
// }

// userSvc 被初始化为 impl.UserServiceImpl 的一个新实例的指针。之后，在其他函数（比如测试函数）中就可以使用这个已经初始化的 userSvc 变量了。
func init() {
	test.DevelopmentSetup()
	userSvc = &impl.UserServiceImpl{}
}
