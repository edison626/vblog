package impl

import (
	"context"

	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/common"
)

func (i *blogServiceImpl) update(ctx context.Context, scope *common.Scope, ins *blog.Blog) error {
	exec := i.db.
		WithContext(ctx).
		Where("id = ?", ins.Id)

	if scope != nil {
		if scope.UserId != "" {
			exec = exec.Where("create_by = ?", scope.UserId)
		}
	}

	return exec.
		Updates(ins).
		Error
}
