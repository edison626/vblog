package controller

import (
	"gitee.com/go-course/go12/skills/mvc/dao"
	"gitee.com/go-course/go12/skills/mvc/model"
)

// 控制器
func CreateUser(ins *model.User) error {

	// 需要Insert 到 数据库里面
	dao.SaveUser(ins)

	return nil
}
