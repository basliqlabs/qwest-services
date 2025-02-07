package userservice

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
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

	// TODO - check for validation errors
	if valid, _ := email.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByEmail(ctx, req.Identifier)
	} else if valid, _ := username.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByUserName(ctx, req.Identifier)
	} else if valid, _ := mobile.IsValid(req.Identifier); valid {
		user, found, err = s.repo.FindUserByMobile(ctx, req.Identifier)
	}

	if err != nil {
		return &userdto.LoginResponse{}, richerror.New(op).WithError(err).WithKind(richerror.KindUnexpected)
	}

	if !found {
		return &userdto.LoginResponse{}, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	areIdentical, err := passwordhash.Compare(user.PasswordHash, req.Password)

	if err != nil {
		return &userdto.LoginResponse{}, richerror.
			New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	if !areIdentical {
		return &userdto.LoginResponse{}, richerror.
			New(op).
			WithKind(richerror.KindNotFound).
			WithMessage(translation.T(lang, "user_not_found"))
	}

	// TODO - fix JWT
	token, err := jwtutil.Generate(user.UserName)
	if err != nil {
		return &userdto.LoginResponse{}, richerror.
			New(op).
			WithKind(richerror.KindUnexpected).
			WithMeta(map[string]any{
				"username": user.UserName,
			})
	}

	return &userdto.LoginResponse{Token: token}, nil
}
