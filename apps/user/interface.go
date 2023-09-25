package user

import (
	"context"

	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 定义User包的能力 就是接口定义
// 站在使用放的角度来定义的   userSvc.Create(ctx, req), userSvc.DeleteUser(id)
// 接口定义好了，不要试图 随意修改接口， 要保证接口的兼容性
type Service interface {
	// 创建用户
	//context.Context: 通常用于跟踪和控制长时间运行的或者需要能够被取消的操作。
	//此方法接收一个 context.Context 和一个指向 CreateUserRequest 的指针，返回一个指向 User 结构体的指针和一个错误（如果有的话）
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
	// 直接序列化为Json存储到 label字段 - https://gorm.io/zh_CN/docs/serializer.html
	Label map[string]string `json:"label" gorm:"serializer:json"`
	// 判断哈希是否被调用
	isHashed bool
}

func (req *CreateUserRequest) SetIsHashed() {
	req.isHashed = true
}

// 校验用户 - 是否为空
func (req *CreateUserRequest) Validate() error {
	if req.Username == "" || req.Password == "" {
		return fmt.Errorf("用户名或者密码需要填写")
	}
	return nil
}

// salt 加盐加密 - 并通过base24的方式存入mysql
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
