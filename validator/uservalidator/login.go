package uservalidator

import (
	"context"

	"github.com/basliqlabs/qwest-services/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/basliqlabs/qwest-services/validator"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx context.Context, req *userdto.LoginRequest) (validator.ValidationErrors, error) {
	lang := contextutil.GetLanguage(ctx)
	const op = "uservalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Identifier,
			validation.Required.Error(translation.TD(lang, "validation.required", map[string]any{
				"Field": translation.T(lang, "fields.identifier"),
			})),
			validation.Length(IdentifierMinLength, IdentifierMaxLength).
				Error(translation.TD(lang, "validation.length", map[string]any{
					"Field":     translation.T(lang, "fields.identifier"),
					"MinLength": IdentifierMinLength,
					"MaxLength": IdentifierMaxLength,
				})),
		),

		validation.Field(&req.Password, validation.Required.Error(translation.TD(lang, "validation.required", map[string]any{
			"Field": translation.T(lang, "fields.password"),
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
