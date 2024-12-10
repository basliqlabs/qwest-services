package validator

import (
	"errors"
	"github.com/basliqlabs/qwest-services-auth/pkg/errmsg"
	"github.com/basliqlabs/qwest-services-auth/pkg/richerror"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ValidationErrors map[string]string

func generateFieldErrors(err error) ValidationErrors {
	fieldErrors := make(ValidationErrors)

	var errV validation.Errors
	ok := errors.As(err, &errV)
	if ok {
		for key, value := range errV {
			if value != nil {
				fieldErrors[key] = value.Error()
			}
		}
	}

	return fieldErrors
}

func GenerateValidationError(op string, req any, err error) (ValidationErrors, error) {
	return generateFieldErrors(err), richerror.
		New(op).
		WithMessage(errmsg.InvalidInput).
		WithKind(richerror.KindInvalid).
		WithMeta(map[string]interface{}{"req": req}).
		WithError(err)
}
