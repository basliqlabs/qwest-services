package authvalidator

import (
	"context"
	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/username"
	"github.com/basliqlabs/qwest-services-auth/validator"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

func (v Validator) Login(ctx *context.Context, req *dto.LoginRequest) (validator.ValidationErrors, error) {
	const op = "authvalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(req.Username,
			validation.Required,
			validation.Match(regexp.MustCompile(username.UsernameRegex)).Error("username is not in correct format"),
		)); err != nil {
		return validator.GenerateValidationError(op, req, err)
	}
	return nil, nil
}
