package user

import "context"

type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
}

type CreateUserRequest struct {
	Username string
	Password string
}
