package userentity

type User struct {
	UserID   int
	UserName string
	Email    string
	Mobile   string
}

type UserWithPasswordHash struct {
	User
	PasswordHash string
}
