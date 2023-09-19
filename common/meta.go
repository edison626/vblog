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
