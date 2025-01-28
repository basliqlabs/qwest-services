package userhandler

import (
	"github.com/basliqlabs/qwest-services/service/userservice"
	"github.com/basliqlabs/qwest-services/validator/uservalidator"
)

type Handler struct {
	validator uservalidator.Validator
	service   userservice.Service
}

func New(validator uservalidator.Validator, service userservice.Service) *Handler {
	return &Handler{validator: validator, service: service}
}
