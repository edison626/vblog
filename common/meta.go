package common

import "time"

func NewMeta() *Meta {
	return &Meta{
		CreatedAt: time.Now().Unix(),
	}
}

type Meta struct {
	// 在添加数据需要村的定义
	Id int64 `json:"id"`
	// 创建时间
	CreatedAt int64 `json:"created_at"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at"`
}

// 控制用户访问数据的访问
// 操作数据的时候, 加上一个where条件
// 比如用户A10, 要去编辑用户B(12)的文章,  id=10 and create_by = 10
type Scope struct {
	UserId string `json:"user_id"`
}
