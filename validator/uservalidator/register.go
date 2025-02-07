package uservalidator

import (
	"context"

	"github.com/basliqlabs/qwest-services/dto/userdto"
	"github.com/basliqlabs/qwest-services/validator"
)

func (v Validator) Register(ctx context.Context, req *userdto.RegisterRequest) (validator.ValidationErrors, error) {
	const op = "uservalidator.Register"

	return nil, nil
}
