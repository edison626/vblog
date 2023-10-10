package user

// 主要定义业务接口与接口参数

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	AppName = "user"
)

// 定义User包的能力 就是接口定义
// 站在使用放的角度来定义的   userSvc.Create(ctx, req), userSvc.DeleteUser(id)
// 接口定义好了，不要试图 随意修改接口， 要保证接口的兼容性
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) error

	// 查询用户  User.CheckPassword(xxx)
	DescribeUserRequest(context.Context, *DescribeUserRequest) (*User, error)
}

func NewDescribeUserRequestById(id string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeValue: id,
	}
}

func NewDescribeUserRequestByUsername(username string) *DescribeUserRequest {
	return &DescribeUserRequest{
		DescribeBy:    DESCRIBE_BY_USERNAME,
		DescribeValue: username,
	}
}

// 同时支持通过Id来查询，也要支持通过username来查询
type DescribeUserRequest struct {
	DescribeBy    DescribeBy `json:"describe_by"`
	DescribeValue string     `json:"describe_value"`
}

func NewCreateUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Role:  ROLE_MEMBER,
		Label: map[string]string{},
	}
}

// VO
type CreateUserRequest struct {
	//
	Username string `json:"username" gorm:"column:username"`
	//
	Password string `json:"password"`
	// 用户的角色(写死)，
	// @Role("admin")
	// CreateBlog( check user Role  )
	Role Role `json:"role"`
	// 对象标签, Dep:部门A
	// Label 没法存入数据库，不是一个结构化的数据
	// 比如就存储在数据里面 ，存储为Json, 需要ORM来帮我们完成 json的序列化和存储
	// 直接序列化为Json存储到 lable字段
	Label map[string]string `json:"label" gorm:"serializer:json"`

	isHashed bool
}

func (req *CreateUserRequest) SetIsHashed() {
	req.isHashed = true
}

func (req *CreateUserRequest) Validate() error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("用户名或者密码需要填写")
	}
	return nil
}

func (req *CreateUserRequest) PasswordHash() {
	if req.isHashed {
		return
	}

	b, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(b)
	req.isHashed = true
}

// 删除用户的请求
type DeleteUserRequest struct {
	Id int `json:"id"`
}

func (req *DeleteUserRequest) IdString() string {
	return fmt.Sprintf("%d", req.Id)
}
