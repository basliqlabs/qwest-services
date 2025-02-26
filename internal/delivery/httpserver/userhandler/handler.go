package userhandler

import (
	"github.com/basliqlabs/qwest-services/internal/delivery/httpserver/middleware"
	"github.com/basliqlabs/qwest-services/internal/service/tokenservice"
	"github.com/basliqlabs/qwest-services/internal/service/userservice"
	"github.com/basliqlabs/qwest-services/internal/validator/uservalidator"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	validator    uservalidator.Validator
	service      userservice.Service
	tokenService tokenservice.Service
	auth         middleware.AuthMiddleware
}

func New(validator uservalidator.Validator, service userservice.Service, tokenService tokenservice.Service, auth middleware.AuthMiddleware) *Handler {
	return &Handler{
		validator:    validator,
		service:      service,
		tokenService: tokenService,
		auth:         auth,
	}
}

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("/login", h.login, h.auth.OptionalAuth())
	userGroup.POST("/register", h.register, h.auth.OptionalAuth())
	userGroup.POST("/logout", h.logout, h.auth.StrictAuth())
}
