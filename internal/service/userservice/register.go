package userservice

import (
	"context"
	"strings"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/internal/entity/userentity"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/passwordhash"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/basliqlabs/qwest-services/pkg/username"
)

func (s *Service) Register(ctx context.Context, req *userdto.RegisterRequest) (*userdto.RegisterResponse, error) {
	const op = "userservice.Register"
	lang := contextutil.GetLanguage(ctx)

	_, exists, _ := s.repo.FindUserByEmail(ctx, req.Email)

	if exists {
		return nil, richerror.New(op).
			WithKind(richerror.KindForbidden).
			WithMessage(translation.T(lang, "user_already_exists"))
	}

	usernameBase := strings.Split(req.Email, "@")[0]

	hashedPassword, err := passwordhash.Hash(req.Password)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	user := userentity.UserWithPasswordHash{
		User: userentity.User{
			Email:    req.Email,
			UserName: username.GenerateUnique(usernameBase),
			Mobile:   "",
		},
		PasswordHash: hashedPassword,
	}

	_, err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, richerror.New(op).
			WithKind(richerror.KindUnexpected).
			WithError(err)
	}

	return nil, nil
}
