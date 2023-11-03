package impl_test

import (
	"testing"

	"github.com/edison626/vblog/apps/blog"
)

func TestCreateBlog(t *testing.T) {
	in := blog.NewCreateBlogRequest()
	in.Title = "Vblog Web Service Api6"
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
	// SELECT * FROM `blogs` WHERE create_by IN ('admin') LIMIT 10
	in.AddUsername("张三")
	in.PageSize = 2
	set, err := svc.QueryBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestDescribeBlog(t *testing.T) {
	in := blog.NewDescribeBlogRequest("45")
	set, err := svc.DescribeBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(set)
}

func TestUpdateBlogPut(t *testing.T) {
	in := blog.NewPutUpdateBlogRequest("59")
	in.Content = "Golang2"
	ins, err := svc.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

// 用Merge 覆盖内容 - 有相同的字段就保留，不一样的就覆盖
func TestUpdateBlogPatch(t *testing.T) {
	in := blog.NewPutUpdateBlogRequest("59")
	in.Title = "Vblog Web Service Api23"
	in.Tags["分类"] = "Golang3"
	ins, err := svc.UpdateBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDeleteBlog(t *testing.T) {
	in := blog.NewDeleteBlogRequest("45")
	err := svc.DeleteBlog(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
}
