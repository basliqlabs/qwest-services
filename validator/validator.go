package validator

import (
	"errors"

	"github.com/basliqlabs/qwest-services-auth/pkg/richerror"
	"github.com/basliqlabs/qwest-services-auth/pkg/translation"
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

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

type Args struct {
	Operation string
	Request   any
	Error     error
}

func (v *Validator) Generate(args Args) (ValidationErrors, error) {
	return generateFieldErrors(args.Error), richerror.
		New(args.Operation).
		WithMessage(translation.T(translation.GetCoreLang(), "invalid_input", nil)).
		WithKind(richerror.KindInvalid).
		WithMeta(map[string]interface{}{"req": args.Request}).
		WithError(args.Error)
}
