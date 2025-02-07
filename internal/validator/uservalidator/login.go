package uservalidator

import (
	"context"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/internal/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx context.Context, req *userdto.LoginRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "uservalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Identifier,
			validator.RequiredRule(lang, "fields.identifier"),
			validator.LengthRule(lang, "fields.identifier", identifierMinLength, identifierMaxLength)),
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
