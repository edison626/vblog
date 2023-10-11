package impl

import (
	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/conf"
	"github.com/edison626/vblog/ioc"
	"gorm.io/gorm"
)

func init() {
	ioc.Controller().Registry(&blogServiceImpl{})
}

type blogServiceImpl struct {
	//db
	db *gorm.DB
}

func (i *blogServiceImpl) Init() error {
	//ioc
	i.db = conf.C().MySQL.GetConn().Debug()
	return nil
}

func (i *blogServiceImpl) Name() string {
	return blog.AppName
}
