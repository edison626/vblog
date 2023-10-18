package impl

import (
	"context"

	"github.com/edison626/vblog/apps/blog"
	"github.com/edison626/vblog/common"
	"github.com/edison626/vblog/exception"
)

func (i *blogServiceImpl) update(ctx context.Context, scope *common.Scope, ins *blog.Blog) error {
	exec := i.db.
		WithContext(ctx).
		Where("id = ?", ins.Id)

	if scope != nil {
		if scope.Username != "" {
			exec = exec.Where("create_by = ?", scope.Username)
		}
	}

	exec = exec.Updates(ins)

	rf := exec.RowsAffected
	if rf == 0 {
		return exception.NewNotFound("blog %d not found", ins.Id)
	}
	return exec.Error
}
