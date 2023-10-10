package user

type User struct {
	Id       int
	CreateAt int64

	//
	*CreateUserRequest
}
