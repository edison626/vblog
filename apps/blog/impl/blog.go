package impl

import (
	"context"
	"errors"

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

// 更新文章
func (i *blogServiceImpl) UpdateBlog(
	ctx context.Context, in *blog.UpdateBlogRequest) (
	*blog.Blog, error) {
	return nil, nil
}

// 修改文章状态
func (i *blogServiceImpl) UpdateBlogStatus(
	ctx context.Context, in *blog.UpdateBlogStatusRequest) (
	*blog.Blog, error) {
	return nil, nil
}

// 删除文章
func (i *blogServiceImpl) DeleteBlog(
	ctx context.Context, in *blog.DeleteBlogRequest) error {
	return nil
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
