package blog

import (
	"context"
)

const (
	AppName = "blogs"
)

// 博客模块接口
type Service interface {
	//创建博客
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	// 查询文章的列表,列表查询，没有必要查询文章具体内容
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	// 详情页面，尽量把相关的数据查询出来，content
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	//修改文章状态
	UpdateBlogStatus(context.Context, *UpdateBlogStatusRequest) (*Blog, error)
	//更新文章
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	//删除文章
	DeleteBlog(context.Context, *DeleteBlogRequest) error
}

func NewDescribeBlogRequest(id string) *DescribeBlogRequest {
	return &DescribeBlogRequest{
		BlogId: id,
	}
}

type DescribeBlogRequest struct {
	BlogId string `json:"blogid"`
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

type BlogSet struct {
	//博客的总数
	Total int64 `json:"total"`
	// 返回的一页的数据
	Items []*Blog `json:"items"`
}

func (s *BlogSet) Add(items ...*Blog) {
	s.Items = append(s.Items, items...)
}

// 页面默认值
func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageSize:   10,
		PageNumber: 1,
	}
}

// 后段分页
type QueryBlogRequest struct {
	//页面大小
	PageSize int `json:"page_size"`
	//当前属于几个页面
	PageNumber int `json:"page_number"`
	// 0 表示草稿状态，要查询所有的博客
	// nil 没有这个过滤条件
	// 0 DRAFT
	// 1 PUBLISHED
	Status *Status `json:"status"`
}

// 依赖数据库，根据分页大小，当前页数可以推到处获取元素的开始和结束位置
// [1,2,3,4,5][6,7,8,9,10] [...]
// limite (offset,limite) limite(0,5)[1,2,3,4,5]
// limite (5*1,5)[6,7,8,9,10]
// limite (5*2,5)[11,12,13,14,15]
func (r *QueryBlogRequest) Offset() int {
	return r.PageSize * (r.PageNumber - 1)
}

func (r *QueryBlogRequest) SetStatus(s Status) {
	r.Status = &s
}

// PublishBlogStatus 不是属于数据库里的结构，所有没有放到model里面
// 接口请求参数的一部分
type UpdateBlogStatusRequest struct {
	// 如果定义一遍文章，使用对象Id，具体的某一篇文章
	BlogId int64 `json:"blog_id"`
	// 修改的状态 ： DRAFT/PUBLISHED
	Status Status `json:"status"`
}

func NewPutUpdateBlogRequest(id string) *UpdateBlogRequest {
	return &UpdateBlogRequest{
		BlogId:            id,
		UpdateMode:        UPDATE_MODE_PUT,
		CreateBlogRequest: NewCreateBlogRequest(),
	}
}

// 区分全量更新/部分更新
type UpdateBlogRequest struct {
	//如果定义一遍文章，使用对象Id，具体的某一篇文章
	BlogId string `json:"blog_id"`
	// 更新方式- 全量更新/部分更新
	UpdateMode UpdateMode `json:"update_mode"`
	// 用户更新请求，用户只传了一个标签
	*CreateBlogRequest
}

func NewDeleteBlogRequest(id string) *DeleteBlogRequest {
	return &DeleteBlogRequest{
		BlogId: id,
	}
}

type DeleteBlogRequest struct {
	//如果定义一遍文章，使用对象Id，具体的某一篇文章
	BlogId string `json:"blog_id"`
}
