package uservalidator

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/internal/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Register(ctx context.Context, req *userdto.RegisterRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "uservalidator.Register"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Email, validator.EmailRule(lang, "fields.email", true)...),
		validation.Field(&req.Password, validator.PasswordRule(lang, "fields.password", true)...,
		)); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
