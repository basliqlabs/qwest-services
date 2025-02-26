package tokenservice

import (
	"context"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
)

func (s *Service) Logout(ctx context.Context, email string) error {
	const op = "tokenservice.Logout"

	lang := contextutil.GetLanguage(ctx)

	refreshToken, err := s.repo.GetByUserEmail(ctx, email)

	if err != nil {
		return richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "token_has_expired"))
	}

	err = s.repo.RevokeByToken(ctx, refreshToken.Token)
	if err != nil {
		return richerror.New(op).
			WithKind(richerror.KindUnauthorized).
			WithMessage(translation.T(lang, "token_has_been_revoked"))
	}

	return nil
}
