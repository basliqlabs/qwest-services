package userhandler

import "github.com/basliqlabs/qwest-services/validator/authvalidator"

type Handler struct {
	validator authvalidator.Validator
}

func New(validator authvalidator.Validator) *Handler {
	return &Handler{validator: validator}
}
