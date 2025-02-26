package tokenservice

import (
	"context"
	"time"

	"github.com/basliqlabs/qwest-services/internal/dto/tokendto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

func (s *Service) RefreshToken(ctx context.Context, req *tokendto.RefreshTokenRequest) (*tokendto.RefreshTokenResponse, error) {
	const op = "userservice.RefreshToken"
	lang := contextutil.GetLanguage(ctx)

	claims, err := jwtutil.Verify(req.RefreshToken, jwtutil.RefreshToken)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithError(err)
	}

	tokenEntity, err := s.repo.GetByToken(ctx, req.RefreshToken)
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

	token, err := jwtutil.Generate(username, email)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	return &tokendto.RefreshTokenResponse{
		AccessToken: token,
	}, nil
}
