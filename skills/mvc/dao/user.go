package dao

import "gitee.com/go-course/go12/skills/mvc/model"

func SaveUser(ins *model.User) error {
	// orm
	// 需要Insert 到 数据库里面
	// db.Save(ins)

	// 自己写SQL
	// insert (xxxx) Values (???)

	return nil
}
