package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/basliqlabs/qwest-services-auth/pkg/envelope"
	"github.com/basliqlabs/qwest-services-auth/service/authservice"
	"github.com/labstack/echo/v4"
)

func (h Handler) userLogin(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(dto.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrBadRequest,
			Message: contextutil.GetTranslation(ctx, "bad_request", nil),
		}))
	}

	// TODO: temporarily using nil as context
	validationErrors, err := h.validator.Login(ctx, req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrInvalidInput,
			Message: contextutil.GetTranslation(ctx, "invalid_input", nil),
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
