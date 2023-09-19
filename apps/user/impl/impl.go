package impl

import (
	"context"

	"github.com/edison626/vblog/apps/user"
	"github.com/edison626/vblog/conf"
	"gorm.io/gorm"
)

// 检查 &UserServiceImpl{} 是否实现了 user.Service 接口
// 显示声明接口实现的语言 都可以 明确约束接口的实现
var _ user.Service = &UserServiceImpl{}

// var _ user.Service = (*UserServiceImpl)(nil)

func NewUserServiceImpl() *UserServiceImpl {
	// db 怎么来?
	// 通过配置 https://gorm.io/zh_CN/docs/index.html
	return &UserServiceImpl{
		db: conf.C().MySQL.GetConn(),
	}
}

type UserServiceImpl struct {
	db *gorm.DB
}

// 创建用户
// 把变量命令 ctx 和 req
func (i *UserServiceImpl) CreateUser(
	ctx context.Context,
	req *user.CreateUserRequest) (
	*user.User, error) {
	// 1. 校验用户参数
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// 2. 生成一个User对象(ORM对象)
	ins := user.NewUser(req)

	// 3. 保存到数据库, ORM怎么知道这个对象保持在那个表里面, 怎么知道行应用如何对应
	// 怎么知道 ins 需要往users表里面存
	// Create  ins.TableName(), orm提供的功能, 看orm的文档
	// gorm:"column:username" 通过struct tag 定义对象列映射关系
	// ctx 取消了? 这个数据会保存吗? 数据库响应慢? 用户取消了操作?
	// i.db.Create(ins)
	// 现在存储在数据库里面的密码时明文的怎么办?
	// 1. 加密(通过密钥可以解密): 1. 对称加密 2. 非对称加密
	// 2. hash(不可逆):   raw---> has str

	// 关于password的使用,  不是 password.decode == req.password
	// 可以比对hash结果:   password.has_code == req.password.has_code
	// hash: md5(x), hmac, sha1, bcrypt: 专门用于密码hash算法(加盐)
	// "golang.org/x/crypto/bcrypt"
	//  123456 --> xxxxxx
	//  12345 ---> xxxx, 123456 ---> yyyy 123456 ---> zzzz
	//  123456 ---> salt.123456 ---> salt.xxxxxxx (结果)
	//  123456(req) ---> salt.123456   ----> salt.xxxxxx
	// if err := i.db.
	// 	WithContext(ctx).
	// 	Create(ins).
	// 	Error; err != nil {
	// 	return nil, err
	// }

	// 4. 返回结果
	return ins, nil
}

// 删除用户
func (i *UserServiceImpl) DeleteUser(
	context.Context, *user.DeleteUserRequest) error {
	return nil
}
