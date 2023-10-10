package impl

import (
	"context"

	"gitee.com/go-course/go12/vblog/apps/user"
	"gitee.com/go-course/go12/vblog/conf"
	"gitee.com/go-course/go12/vblog/exception"
	"gitee.com/go-course/go12/vblog/ioc"
	"gorm.io/gorm"
)

// 导入这个包的时候, 直接把这个对象 UserServiceImpl 注册给Ioc
//
//	UserServiceImpl
//
// 注册User业务模块(业务模块的名称是user.AppName)的控制器
// User Service 的具体实现 &UserServiceImpl{} 注入
// 可以随时更换业务的具体实现

// 对象的初始化? 删除? .....
// &UserServiceImpl{
// 开启Debug模式
//
//		db: conf.C().MySQL.GetConn().Debug(),
//	}
//
// 提出来 放到初始化里面
func init() {
	ioc.Controller().Registry(&UserServiceImpl{})
}

// 显示声明接口实现的语言 都可以 明确约束接口的实现
var _ user.Service = &UserServiceImpl{}

// var _ user.Service = (*UserServiceImpl)(nil)

// db 怎么来?
// 通过配置 https://gorm.io/zh_CN/docs/index.html
// 定义对象的初始化
func (i *UserServiceImpl) Init() error {
	i.db = conf.C().MySQL.GetConn().Debug()
	return nil
}

// 定义托管到Ioc里面的名称
func (i *UserServiceImpl) Name() string {
	return user.AppName
}

// 他是user service 服务的控制器
type UserServiceImpl struct {
	db *gorm.DB
}

// 创建用户
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
	if err := i.db.
		WithContext(ctx).
		Create(ins).
		Error; err != nil {
		return nil, err
	}

	// 4. 返回结果
	return ins, nil
}

// 删除用户
func (i *UserServiceImpl) DeleteUser(
	ctx context.Context,
	req *user.DeleteUserRequest) error {
	_, err := i.DescribeUserRequest(ctx,
		user.NewDescribeUserRequestById(req.IdString()))
	if err != nil {
		return err
	}

	return i.db.
		WithContext(ctx).
		Where("id = ?", req.Id).
		Delete(&user.User{}).
		Error
}

// 怎么查询一个用户
func (i *UserServiceImpl) DescribeUserRequest(
	ctx context.Context,
	req *user.DescribeUserRequest) (
	*user.User, error) {

	query := i.db.WithContext(ctx)

	// 1. 构造我们的查询条件
	// 根据条件来构建Where语句
	// id= ? or username = ?
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		// 通过返回值来修改原来的对象
		query = query.Where("id = ?", req.DescribeValue)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("username = ?", req.DescribeValue)
	}

	// SELECT * FROM `users` WHERE username = 'admin' ORDER BY `users`.`id` LIMIT 1
	ins := user.NewUser(user.NewCreateUserRequest())
	if err := query.First(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user %s not found", req.DescribeValue)
		}
		return nil, err
	}

	// 数据库里面存储的就是Hash
	ins.SetIsHashed()

	return ins, nil
}
