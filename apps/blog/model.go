package blog

import (
	"encoding/json"
	"time"
)

func NewCreateBlogRequest() *CreateBlogRequest {
	return &CreateBlogRequest{
		Tags: map[string]string{},
	}
}

// 用户参数
type CreateBlogRequest struct {
	//文章标题
	Title string `json:"title"`
	//作者
	Author string `json:"author"`
	// 用户登陆后，我们通过Token 知道是哪个用户
	CreateBy string `json:"create_by"`
	//文章内容
	Content string `json:"content"`
	//概要
	Summary string `json:"summary"`
	//标签，基于标签做分类，语言：golang
	Tags map[string]string `json:"tags" gorm:"serializer:json"`
}

func NewBlog(req *CreateBlogRequest) *Blog {
	return &Blog{
		CreatedAt:         time.Now().Unix(),
		Status:            STATUS_DRAFT,
		CreateBlogRequest: req,
	}
}

// 程序用 - 不可变更
type Blog struct {
	//文章的唯一标识符，给程序使用
	Id int64 `json:"id"`

	//创建时间
	CreatedAt int64 `json:"created_at"`
	//更新时间
	UpdatedAt int64 `json:"updated_at"`
	//发布时间
	PublishedAt int64 `json:"published_at"`
	// 文章的状态
	Status Status `json:"status" `
	// 审核时间
	AuditAt int64 `json:"audit_at"`
	//是否审核成功
	IsAuditPass bool `json:"is_audit_pass"`
	//用户创建博客参数
	*CreateBlogRequest
}

func (b *Blog) TableName() string {
	return "blogs"
}

func (b *Blog) String() string {
	dj, _ := json.Marshal(b)
	return string(dj)
}
