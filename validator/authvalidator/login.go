package authvalidator

import (
	"context"
	"fmt"
	"regexp"

	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/username"
	"github.com/basliqlabs/qwest-services-auth/validator"
	_ "github.com/go-ozzo/ozzo-validation/v4"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (v Validator) Login(ctx *context.Context, req *dto.LoginRequest) (validator.ValidationErrors, error) {
	const op = "authvalidator.Login"

	if err := validation.ValidateStruct(req,
		validation.Field(&req.Username,
			validation.Required,
			validation.Match(regexp.
				MustCompile(username.UsernameRegex)).
				Error(v.util.Translate.T(
					"en",
					"username is not in correct format",
					nil,
				)),
		)); err != nil {
		fmt.Println("####form error", err)
		return v.util.Generate(validator.Args{
			Request: req,
			// TODO: should get this from context
			Language:  "en",
			Operation: op,
			Error:     err,
		})
	}
	return nil, nil
}
