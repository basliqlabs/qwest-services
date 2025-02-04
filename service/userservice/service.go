package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/entity/userentity"
)

type Repository interface {
	FindUserByMobile(ctx context.Context, mobile string) (userentity.UserWithPasswordHash, error)
	FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, error)
	FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{repo: repo}
}
