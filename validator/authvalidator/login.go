package authvalidator

import (
	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/validator"
)

func Login(user *dto.LoginRequest) (validator.ValidationErrors, error) {
	// TODO: implement validation
	return nil, nil
}
