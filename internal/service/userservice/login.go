package userservice

import (
	"context"
	"time"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/internal/entity/tokenentity"
	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/email"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/mobile"
	"github.com/basliqlabs/qwest-services/pkg/passwordhash"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/basliqlabs/qwest-services/pkg/username"
)

func (s *Service) Login(ctx context.Context, req *userdto.LoginRequest) (*userdto.LoginResponse, error) {
	const op = "userservice.Login"
	lang := contextutil.GetLanguage(ctx)

	var (
		user  userentity.UserWithPasswordHash
		found       = false
		err   error = nil
	)

	if valid, _ := email.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByEmail(ctx, req.Identifier)
	} else if valid, _ := username.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByUserName(ctx, req.Identifier)
	} else if valid, _ := mobile.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByMobile(ctx, req.Identifier)
	}

	if err != nil {
		return nil, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if !found {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	areIdentical, err := passwordhash.Compare(user.PasswordHash, req.Password)

	if err != nil {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	if !areIdentical {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	tokenPair, err := jwtutil.GenerateTokenPair(user.UserName, user.Email)
	if err != nil {
		return nil, richerror.
			New(op).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]any{
				"username": user.UserName,
			})
	}

	refreshToken := tokenentity.RefreshToken{
		UserID:    user.UserID,
		Token:     tokenPair.RefreshToken,
		ExpiresAt: time.Now().Add(s.authConfig.JWT.RefreshTokenExpirationTime),
		CreatedAt: time.Now(),
		Revoked:   false,
	}

	err = s.tokenRepo.Create(ctx, refreshToken)
	if err != nil {
		return &userdto.LoginResponse{}, richerror.
			New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	return &userdto.LoginResponse{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
