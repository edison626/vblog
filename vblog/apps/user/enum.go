package user

// 使用这个ROLE类型来表现枚举类型
type Role int

const (
	// 普通用户
	ROLE_MEMBER Role = iota
	// 管理员
	ROLE_ADMIN
)

type DescribeBy int

const (
	DESCRIBE_BY_ID DescribeBy = iota
	DESCRIBE_BY_USERNAME
)
