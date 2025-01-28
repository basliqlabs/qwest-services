package userservice

import "context"

type Repository interface {
	DoesUserNameWithPasswordExist(ctx context.Context, username string, password string) (bool, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
