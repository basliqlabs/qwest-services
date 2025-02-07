package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
)

func (s *Service) Register(ctx context.Context, req *userdto.RegisterRequest) (*userdto.RegisterResponse, error) {
	const op = "userservice.Register"
	return nil, nil
}
