package user_test

import (
	"context"
	"testing"

	"gitee.com/go-course/go12/skills/ddd/user"
)

func TestCreateUserIF(t *testing.T) {
	// 接口是个规范 svc为什么为nil
	var svc user.Service

	svc = &UserServceImpl{}

	svc.CreateUser(
		context.Background(),
		&user.CreateUserRequest{},
	)
}

// 是一个Mock实现(测试用例里面使用)
type UserServceImpl struct {
}

func (i *UserServceImpl) CreateUser(
	context.Context, *user.CreateUserRequest) (
	*user.User,
	error,
) {
	// 具体怎么做
	return &user.User{}, nil
}
