package uservalidator

import "github.com/basliqlabs/qwest-services/validator"

type Validator struct {
	util *validator.Validator
}

func New(util *validator.Validator) Validator {
	return Validator{
		util: util,
	}
}

const IdentifierMinLength = 6
const IdentifierMaxLength = 100
