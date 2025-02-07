package userhandler

import (
	"github.com/basliqlabs/qwest-services/service/userservice"
	"github.com/basliqlabs/qwest-services/validator/uservalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator uservalidator.Validator
	service   userservice.Service
}

func New(validator uservalidator.Validator, service userservice.Service) *Handler {
	return &Handler{validator: validator, service: service}
}

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/login", h.login)
	userGroup.POST("/register", h.register)
}
