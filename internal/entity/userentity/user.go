package userentity

type User struct {
	UserName string
	Email    string
	Mobile   string
}

type UserWithPasswordHash struct {
	User
	PasswordHash string
}
