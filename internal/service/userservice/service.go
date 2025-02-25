package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/config"
	"github.com/basliqlabs/qwest-services/internal/entity/tokenentity"
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

type RefreshTokenRepository interface {
	Create(ctx context.Context, token tokenentity.RefreshToken) (tokenentity.RefreshToken, error)
	GetByToken(ctx context.Context, token string) (tokenentity.RefreshToken, error)
	GetByUserEmail(ctx context.Context, email string) (tokenentity.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID int) error
	DeleteByToken(ctx context.Context, token string) error
	RevokeByToken(ctx context.Context, token string) error
}

type Service struct {
	repo       Repository
	tokenRepo  RefreshTokenRepository
	jwt        jwtutil.JWT
	authConfig config.AuthConfig
}

func New(repo Repository, tokenRepo RefreshTokenRepository, jwt jwtutil.JWT, authConfig config.AuthConfig) Service {
	return Service{
		repo:       repo,
		tokenRepo:  tokenRepo,
		jwt:        jwt,
		authConfig: authConfig,
	}
}
