package impl

import (
	"context"
	"errors"
	"fmt"
	"time"

	"dario.cat/mergo"
	"github.com/edison626/vblog/apps/blog"
)

// 创建博客
func (i *blogServiceImpl) CreateBlog(
	ctx context.Context, in *blog.CreateBlogRequest) (
	*blog.Blog, error) {

	ins := blog.NewBlog(in)
	//.... 业务逻辑
	if i.db == nil {
		return nil, errors.New("database connection is not initialized")
	}
	if err := i.db.WithContext(ctx).Create(ins).Error; err != nil {
		return nil, err
	}
	return ins, nil
}

func (i *blogServiceImpl) QueryBlog(
	ctx context.Context, in *blog.QueryBlogRequest) (
	*blog.BlogSet, error) {
	query := i.db.WithContext(ctx).Model(&blog.Blog{})

	//提前准备好Set对象
	set := blog.NewBlogSet()

	// 组装查询条件
	if in.Status != nil {
		query = query.Where("Status = ?", *in.Status)
	}
	if in.Keywords != "" {
		// 关键字过滤: 模糊匹配,  Golang入门   Golang
		// ? 占位符 '%Test12%'
		query = query.Where("title LIKE ?", "%"+in.Keywords+"%")
	}
	if len(in.Usernames) > 0 {
		// gorm 会把 [] --> (xxx,xxx)
		query = query.Where("create_by IN ?", in.Usernames)
	}

	// 1. 查询总数量 - 参考 https://arco.design/vue/component/pagination
	err := query.Count(&set.Total).Error
	if err != nil {
		return nil, err
	}

	// 2. 查询一页的数据
	err = query.
		Offset(in.Offset()).
		Limit(in.PageSize).
		Find(&set.Items).
		Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

// 详情页面，尽量把相关的数据查询出来，content
func (i *blogServiceImpl) DescribeBlog(
	ctx context.Context, in *blog.DescribeBlogRequest) (
	*blog.Blog, error) {

	query := i.db.WithContext(ctx).Model(&blog.Blog{})
	ins := blog.NewBlog(blog.NewCreateBlogRequest())

	err := query.Where("id = ?", in.BlogId).First(ins).Error
	if err != nil {
		return nil, err
	}
	return ins, nil

}

// 更新文章
func (i *blogServiceImpl) UpdateBlog(
	ctx context.Context, in *blog.UpdateBlogRequest) (
	*blog.Blog, error) {

	//查询更新对象
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}

	//
	switch in.UpdateMode {
	case blog.UPDATE_MODE_PUT:
		//全量更新
		ins.CreateBlogRequest = in.CreateBlogRequest
	case blog.UPDATE_MODE_PATCH:
		//增量更新
		// if in.Author != "" {
		// 	ins.Author = in.Author
		// }
		// if in.Content != "" {
		// 	ins.Content = in.Content
		// }

		//有没有工具来帮我们完成2个结构题的merge
		err := mergo.Merge(
			ins.CreateBlogRequest,
			in.CreateBlogRequest,
			mergo.WithOverride)
		if err != nil {
			return nil, err
		}

	default:
		return nil, fmt.Errorf("unknown update mode: %d", in.UpdateMode)
	}

	//更新数据
	//更新的sql 命令 - UPDATE `blogs` SET `created_at`=1697010768,`updated_at`=1697094685,`status`='1',`title`='Vblog Web Service Api23',`tags`='{"分类":"Golang3"}' WHERE id = '45' AND `id` = 45
	ins.UpdatedAt = time.Now().Unix()
	//err = i.db.WithContext(ctx).Where("id = ?", in.BlogId).Updates(ins).Error
	err = i.update(ctx, in.Scope, ins)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// 文章审核, 审核通过的文章才能被看到
func (i *blogServiceImpl) AuditBlog(
	ctx context.Context, in *blog.AuditBlogRequest) (
	*blog.Blog, error) {
	// 查询需要更新的对象
	ins, err := i.DescribeBlog(ctx, blog.NewDescribeBlogRequest(in.BlogId))
	if err != nil {
		return nil, err
	}

	ins.IsAuditPass = in.IsAuditPass
	ins.AuditAt = time.Now().Unix()
	err = i.update(ctx, nil, ins)
	if err != nil {
		return nil, err
	}
	return ins, err
}

// 删除文章
func (i *blogServiceImpl) DeleteBlog(
	ctx context.Context, in *blog.DeleteBlogRequest) error {
	//执行sql 命令 - DELETE FROM `blogs` WHERE id = '45'
	return i.db.WithContext(ctx).
		Model(&blog.Blog{}).
		Where("id = ?", in.BlogId).
		Delete(&blog.Blog{}).
		Error
}

// 修改文章状态
func (i *blogServiceImpl) UpdateBlogStatus(
	ctx context.Context, in *blog.UpdateBlogStatusRequest) (
	*blog.Blog, error) {
	return nil, nil
}
