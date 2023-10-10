package user

// 用于定于定义对象模型(存入数据库里面的对象)

import (
	"encoding/json"

	"gitee.com/go-course/go12/vblog/common"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(req *CreateUserRequest) *User {
	req.PasswordHash()

	return &User{
		Meta:              common.NewMeta(),
		CreateUserRequest: req,
	}
}

// 用于存放 存入数据库的对象(PO)
type User struct {
	// 通用信息
	*common.Meta
	// 用户传递过来的请求
	*CreateUserRequest
}

func (u *User) String() string {
	dj, _ := json.Marshal(u)
	return string(dj)
}

// 判断该用户的密码是否正确
func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// 声明你这个对象存储在users表里面
// orm 负责调用TableName() 来动态获取你这个对象要存储的表的名称
func (u *User) TableName() string {
	return "users"
}
