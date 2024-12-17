package authvalidator

import (
	"context"
	"regexp"

	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/basliqlabs/qwest-services-auth/pkg/username"
	"github.com/basliqlabs/qwest-services-auth/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx context.Context, req *dto.LoginRequest) (validator.ValidationErrors, error) {
	coreLang := contextutil.GetCoreLang(ctx)
	const op = "authvalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Username,
			validation.Required.Error(contextutil.GetTranslation(ctx, "validation.required", map[string]any{
				"Field": contextutil.GetTranslation(ctx, "fields.username", nil),
			})),
			validation.Match(regexp.
				MustCompile(username.UsernameRegex)).
				Error(contextutil.GetTranslation(ctx, "validation.invalid", map[string]any{
					"Field": contextutil.GetTranslation(ctx, "fields.username", nil),
				})),
		)); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Language:  coreLang,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
