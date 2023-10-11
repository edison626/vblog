package impl_test

import (
	"context"

	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/ioc"
	"github.com/edison626/vblog/test"
)

var (
	svc blog.Service
	ctx = context.Background()
)

func init() {
	test.DevelopmentSetup()
	svc = ioc.Controller().Get(blog.AppName).(blog.Service)
}
