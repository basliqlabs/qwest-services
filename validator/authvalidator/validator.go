package authvalidator

import "github.com/basliqlabs/qwest-services-auth/validator"

type Validator struct {
	util validator.Validator
}

func New(util validator.Validator) Validator {
	return Validator{
		util: util,
	}
}
