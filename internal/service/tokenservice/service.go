package tokenservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/config"
	"github.com/basliqlabs/qwest-services/internal/entity/tokenentity"
)

type Repository interface {
	Create(ctx context.Context, token tokenentity.RefreshToken) error
	GetByToken(ctx context.Context, token string) (tokenentity.RefreshToken, error)
	GetByUserEmail(ctx context.Context, email string) (tokenentity.RefreshToken, error)
	DeleteByUserID(ctx context.Context, userID int) error
	DeleteByToken(ctx context.Context, token string) error
	RevokeByToken(ctx context.Context, token string) error
}

type Service struct {
	repo       Repository
	authConfig config.AuthConfig
}

func New(repo Repository, authConfig config.AuthConfig) Service {
	return Service{
		repo:       repo,
		authConfig: authConfig,
	}
}
