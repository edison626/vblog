package blog

type Status int

const (
	// 文章状态草稿
	STATUS_DRAFT Status = iota
	// 已经发布
	STATUS_PUBLISHED
)

type UpdateMode int

const (
	//全量更新
	UPDATE_MODE_PUT UpdateMode = iota
	// 部分更新（增量更新）
	UPDATE_MODE_PATCH
)
