package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/dto/userdto"
)

func (s *Service) Login(ctx context.Context, req *userdto.LoginRequest) (*userdto.LoginResponse, error) {
	_, err := s.repo.DoesUserNameWithPasswordExist(ctx, req.Identifier, req.Password)
	return nil, err
}
