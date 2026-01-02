package user

type User struct {
}

func NewUser() IUser {
	return &User{}
}
