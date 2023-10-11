package impl_test

import (
	"testing"

	"github.com/edison626/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Title = "Vblog Web Service Api3"
	in.Content = "Golong "
	in.Tags["分类"] = "Golang"
	ins, err := svc.CreateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	in := blog.NewQueryBlogRequest()
	in.SetStatus(blog.STATUS_PUBLISHED)
	set, err := svc.QueryBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}
