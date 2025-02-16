package userentity

type User struct {
	UserName string
	Email    string
}

type UserWithPasswordHash struct {
	User
	PasswordHash string
}
