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
	t.Log(req)
	u, err := userSvc.CreateUser(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)

}

// func TestDeleteUser(t *testing.T) {
// 	err := userSvc.DeleteUser(ctx, &user.DeleteUserRequest{
// 		//ID: 8, //删除ID8
// 	})
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// }

// 测试查询 - name 查询
// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
// func TestDescribeUserRequestByName(t *testing.T) {
// 	req := user.NewDescribeUserRequestByUsername("admin")
// 	ins, err := userSvc.DescribeUserRequest(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(ins)
// }

// 测试查询 - ID 查询
// 修改ID 和 mysql 的id 一样
// func TestDescribeUserRequestById(t *testing.T) {
// 	req := user.NewDescribeUserRequestById("9")
// 	ins, err := userSvc.DescribeUserRequest(ctx, req)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	t.Log(ins)
// }

// userSvc 被初始化为 impl.UserServiceImpl 的一个新实例的指针。之后，在其他函数（比如测试函数）中就可以使用这个已经初始化的 userSvc 变量了。
func init() {
	//conf.Setup("")
	test.DevelopmentSetup()
	userSvc = &impl.UserServiceImpl{}
}
