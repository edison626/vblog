package api

import (
	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/ioc"
)

func init() {
	ioc.ApiHandler().Registry(&apiHandler{})
}

type apiHandler struct {
	svc blog.Service
}

func (t *apiHandler) Name() string {
	return blog.AppName
}

func (t *apiHandler) Init() error {
	t.svc = ioc.Controller().Get(blog.AppName).(blog.Service)
	return nil
}
