package user

// 使用这个ROLE类型来表现枚举类型 iota
type Role int

const (
	// 创建者，负责写博客的 - 参数值=0
	ROLE_AUTHOR Role = iota
	// 审核人员 - 参数值=1
	ROLE_AUDITOR
	//管理员 - 参数值=2
	ROLE_ADMIN
)

type DescribeBy int

const (
	DESCRIBE_BY_ID DescribeBy = iota
	DESCRIBE_BY_USERNAME
)
