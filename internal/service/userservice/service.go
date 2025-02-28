package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
)

type Repository interface {
	// identifierType can be "username", "email", or "mobile"
	DoesIdentifierWithPasswordExist(ctx context.Context, identifier string, password string, identifierType string) (bool, error)
	FindUserByMobile(ctx context.Context, mobile string) (userentity.UserWithPasswordHash, bool, error)
	FindUserByEmail(ctx context.Context, email string) (userentity.UserWithPasswordHash, bool, error)
	FindUserByUserName(ctx context.Context, username string) (userentity.UserWithPasswordHash, bool, error)
	CreateUser(ctx context.Context, user userentity.UserWithPasswordHash) (int, error)
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}
