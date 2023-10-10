# 令牌管理


## Restful设计接口

+ http method: POST/GET/PUT/PATCH/..
+ http path:   /users/list,  /users/get

HTTP API:  通过HTTP协议把请求传递给你服务端, 服务端详情请求

传统的API风格:
Http Method: POST 
Http URL：  /object/action,  /users/list /users/create /users/login
Http Body: 存放参数, url encod 编码

Restful 风格:
API: 资源的状态变化:  REST即表述性状态传递（英文：Representational State Transfer，简称REST

定义一种操作服务端资源的一种API风格

资产: users, blogs, 
动作: 使用http method: POST:创建, GET获取, PUT:全量修改, PATCH:修改部分属性, DELETE: 删除

POST: /users/ : 创建用户
GET:  /usres/ : 获取用于列表
GET: /users/01/: 获取id为01的资源详情
PUT: /usrs/01/: 修改id为01的user资源
DELETE /users/01/: 删除id为01的资源


令牌管理的RestulfAPI

+ POST /tokens/: 颁发token, 参数在HTTP Body
+ DELETE /tokens/: 删除token, 参数在Header
+ 统一采用JSON来交互: JSON ON HTTP

接口改进:  /xxxx/xxx ---> 静态服务器(html),   /api ---> localhost:8070

// 代表这个是个后端接口
/api/service_name/service_versions
+ POST /api/vblog/v1/tokens/: 颁发token, 参数在HTTP Body
+ DELETE /api/vblog/v1/tokens/: 删除token, 参数在Header

通过HTTP协议传递参数的方式:
+ 1. 通过Header:  AccessToken: xxxxxx
+ 2. URL
   + 路径参数: /users/01  01就是路径参数
   + URL参数: /users/01?key1=value1&key2=values   key1=value1&key2=values url参数
+ 3. Body
   + 通过HTTP Body来传递参数


