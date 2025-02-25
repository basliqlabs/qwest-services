package userservice

import (
	"context"
	"time"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

func (s *Service) RefreshToken(ctx context.Context, req *userdto.RefreshTokenRequest) (*userdto.RefreshTokenResponse, error) {
	const op = "userservice.RefreshToken"
	lang := contextutil.GetLanguage(ctx)

	claims, err := s.jwt.Verify(req.RefreshToken, jwtutil.RefreshToken)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithError(err)
	}

	tokenEntity, err := s.tokenRepo.GetByToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindNotFound).
			WithError(err)
	}

	if tokenEntity.Revoked {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "token_has_been_revoked"))
	}

	if time.Now().After(tokenEntity.ExpiresAt) {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "token_has_expired"))
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "invalid_token_claims"))
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "invalid_token_claims"))
	}

	accessToken, err := s.jwt.Generate(username, email)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	return &userdto.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

func (s *Service) Logout(ctx context.Context, req *userdto.LogoutRequest) error {
	const op = "userservice.Logout"

	_, err := s.jwt.Verify(req.RefreshToken, jwtutil.RefreshToken)
	if err != nil {
		return richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithError(err)
	}

	err = s.tokenRepo.RevokeByToken(ctx, req.RefreshToken)
	if err != nil {
		return richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	return nil
}
