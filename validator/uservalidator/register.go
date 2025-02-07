package uservalidator

import (
	"context"

	"github.com/basliqlabs/qwest-services/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/email"
	"github.com/basliqlabs/qwest-services/pkg/password"
	"github.com/basliqlabs/qwest-services/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Register(ctx context.Context, req *userdto.RegisterRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "uservalidator.Register"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Email,
			validator.RequiredRule(lang, "fields.email"),
			validator.LengthRule(lang, "fields.email", email.MinLength, email.MaxLength),
			validator.FormatRule(lang, "validation.email", email.Regex)),
		validation.Field(&req.Password,
			validator.RequiredRule(lang, "fields.password"),
			validator.LengthRule(lang, "fields.password", password.MinLength, password.MaxLength),
			validator.FormatRule(lang, "fields.password", password.Regex),
		)); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
