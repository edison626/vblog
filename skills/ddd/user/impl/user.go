package impl

import (
	"context"

	"gitee.com/go-course/go12/skills/ddd/user"
)

var (
	// 手动补充约束条件: Java 显示接口
	// impl user.Service = &Impl{}

	// int(x) (int)(x)  (*int)(nil)
	// 就是一个约束条件, 不会产生内存, 编译器帮你做类型检查
	_ user.Service = (*Impl)(nil)
)

func NewImpl() *Impl {
	return &Impl{}
}

// impl 他就是来实现 User Service 业务
type Impl struct {
}

func (i *Impl) CreateUser(
	context.Context, *user.CreateUserRequest) (
	*user.User,
	error,
) {
	// 具体怎么做
	return nil, nil
}
