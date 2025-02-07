package uservalidator

import (
	"github.com/basliqlabs/qwest-services/pkg/email"
	"github.com/basliqlabs/qwest-services/pkg/username"
	"github.com/basliqlabs/qwest-services/validator"
)

type Validator struct {
	util *validator.Validator
}

func New(util *validator.Validator) Validator {
	return Validator{
		util: util,
	}
}

const identifierMinLength = username.MinUserNameLength
const identifierMaxLength = email.MaxLength
