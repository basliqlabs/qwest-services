package authvalidator

import (
	"context"
	"regexp"

	"github.com/basliqlabs/qwest-services/dto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/basliqlabs/qwest-services/pkg/username"
	"github.com/basliqlabs/qwest-services/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx context.Context, req *dto.LoginRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "authvalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Username,
			validation.Required.Error(translation.T(lang, "validation.required", map[string]any{
				"Field": translation.T(lang, "fields.username", nil),
			})),
			validation.Match(regexp.MustCompile(username.UsernameRegex)).
				Error(translation.T(lang, "validation.invalid", map[string]any{
					"Field": translation.T(lang, "fields.username", nil),
				})),
		)); err != nil {
		return v.util.Generate(validator.Args{
			Request:   req,
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
