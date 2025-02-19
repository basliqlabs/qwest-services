package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
)

type Repository interface {
	FindUserByMobile(ctx context.Context, mobile string) (userentity.UserWithPasswordHash, bool, error)
	FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, bool, error)
	FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, bool, error)
	DoesUserNameWithPasswordExist(ctx context.Context, username string, password string) (bool, error)
	CreateUser(ctx context.Context, user userentity.UserWithPasswordHash) (int, error)
}

type Service struct {
	repo Repository
	jwt  jwtutil.JWT
}

func New(repo Repository, jwt jwtutil.JWT) Service {
	return Service{
		repo: repo,
		jwt:  jwt,
	}
}
