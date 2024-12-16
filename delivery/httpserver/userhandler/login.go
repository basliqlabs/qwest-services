package userhandler

import (
	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/envelope"
	"github.com/basliqlabs/qwest-services-auth/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {
	req := new(dto.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrBadRequest,
			Message: "Message for bad request bla bla bla",
		}))
	}

	// TODO: temporarily using nil as context
	validationErrors, err := h.validator.Login(nil, req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrInvalidInput,
			Message: "invalid form input...",
			Fields:  validationErrors,
		}))
	}

	// RESEARCH: pointer vs concrete structs
	authSvc := authservice.New("")
	res, err := authSvc.Login(req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, envelope.FromRichError(err))
	}

	return c.JSON(http.StatusOK, envelope.New(true).WithData(map[string]any{
		"login": res,
	}))
}
