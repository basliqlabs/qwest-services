package validator

import (
	"errors"
	"github.com/basliqlabs/qwest-services-auth/pkg/errmsg"
	"github.com/basliqlabs/qwest-services-auth/pkg/richerror"
	"github.com/basliqlabs/qwest-services-auth/translation"
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
	Translate translation.Translator
	err       errmsg.ErrorMessage
}

func New(t translation.Translator) *Validator {
	e := errmsg.New(t)
	return &Validator{
		Translate: t,
		err:       e,
	}
}

type Args struct {
	Language  string
	Operation string
	Request   any
	Error     error
}

func (v *Validator) Generate(args Args) (ValidationErrors, error) {
	return generateFieldErrors(args.Error), richerror.
		New(args.Operation).
		WithMessage(v.err.InvalidInput(args.Language)).
		WithKind(richerror.KindInvalid).
		WithMeta(map[string]interface{}{"req": args.Request}).
		WithError(args.Error)
}
