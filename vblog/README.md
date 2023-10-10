# vblog 微博客项目

## 需求

想要Markdown来写一本书 gitbook/typora

Markdown博客的需求

## 产品原型

用户: 
  + 访客：不用登陆 就能浏览文章, 登录后才能进行评论
  + 博客写手: 创作者登录后, 发布文章(Markdown编辑器) 

流程: 发布博客, 访客可以在界面搜索并且查看博客

产品原型:

[](./docs/proto.drawio)

## 软件架构

[](./docs/arch.drawio)

+ api server
+ ui

业务功能架构架构
[](./docs/featrue.drawio)

## 产品的研发

### 后端开发

1. 接口设计

[](./docs/interface.drawio)

2. 流程设计

做为一个项目过程, 他有哪些部件构建(项目骨架)

[](./docs/apiserver.drawio)

统一采用 轻量级的DDD模式

3. 对象与数据
[](./docs/data_controller.drawio)


vblog项目初始化
```
go mod init gitee.com/go-course/go12/vblog
```

4. 编程方式

[](./docs/program.drawio)

#### 编写user模块

1. Interface定义
```go
// 定义User包的能力 就是接口定义
// 站在使用放的角度来定义的   userSvc.Create(ctx, req), userSvc.DeleteUser(id)
// 接口定义好了，不要试图 随意修改接口， 要保证接口的兼容性
type Service interface {
	// 创建用户
	CreateUser(context.Context, *CreateUserRequest) (*User, error)
	// 删除用户
	DeleteUser(context.Context, *DeleteUserRequest) error
}
```
2. Interface实现(TDD)

+ 2.1 定义一个对象来实现这个接口
+ 2.2 补充依赖的配置管理
+ 2.3 单元测试如何读取到配置
+ 2.4 补充数据库的表
+ 2.5 程序当中的异常如此处理? 都通过Error返回吗?

#### 编写token模块

+ 2.1 定义一个对象来实现这个接口
+ 2.2 实现接口(TDD)

+ 2.3 Restful接口开发
Restful API(Web Service) 基于Gin框架处理HTTP协议数据
```go
// Login HandleFunc
func (h *TokenApiHandler) Login(c *gin.Context) {
	// 1. 获取用户的请求参数， 参数在Body里面
	// 一定要使用JSON
	req := token.NewLoginRequest()

	// json.unmarsal
	// http boyd ---> LoginRequest Object
	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 2. 执行逻辑
	// 把http 协议的请求 ---> 控制器的请求
	ins, err := h.svc.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// 3. 返回响应
	c.JSON(http.StatusOK, ins)
}
```


```go
// 需要把HandleFunc 添加到Root路由，定义 API ---> HandleFunc
// 可以选择把这个Handler上的HandleFunc都注册到路由上面
func (h *TokenApiHandler) Registry(r gin.IRouter) {
	// r 是Gin的路由器
	r.POST("/tokens/", h.Login)
	r.DELETE("/tokens/", h.Logout)
}
```

#### 程序入口

把我们控制器和对象组织启动, 启动程序


#### 接口的数据格式统计

```go
// 正常请求数据返回
func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// 异常情况的数据返回, 返回我们的业务Exception
func Failed(c *gin.Context, err error) {
	var e *exception.ApiException
	if v, ok := err.(*exception.ApiException); ok {
		e = v
	} else {
		// 非可以预期, 没有定义业务的情况
		e = exception.New(http.StatusInternalServerError, err.Error())
		e.HttpCode = http.StatusInternalServerError
	}

	c.JSON(e.HttpCode, e)
}
```

第一个小版本

#### 如何优雅的解决对象依赖

```go

//2. 初始化控制
// 2.1 user controller
userServiceImpl := userImpl.NewUserServiceImpl()

// 2.2 token controller
tokenServiceImpl := tokenImpl.NewTokenServiceImpl(userServiceImpl)

// 2.3 token api handler
tkApiHandler := tokenApiHandler.NewTokenApiHandler(tokenServiceImpl)

// ...
```