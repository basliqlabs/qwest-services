package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
)

// userLogin lets a user get an authorization token
//
//	@Summary		User login
//	@Description	Authenticate a user with username and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		userdto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=userdto.LoginResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/login [post]
func (h Handler) userLogin(c echo.Context) error {
	ctx := c.Request().Context()
	lang := contextutil.GetLanguage(ctx)
	req := new(userdto.LoginRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrBadRequest,
			Message: translation.T(lang, "bad_request", nil),
		}))
	}

	validationErrors, err := h.validator.Login(ctx, req)

	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, envelope.New(false).WithError(&envelope.ResponseError{
			Code:    envelope.ErrInvalidInput,
			Message: translation.T(lang, "invalid_input", nil),
			Fields:  validationErrors,
		}))
	}

	res, err := h.service.Login(ctx, req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, envelope.FromRichError(err))
	}

	return c.JSON(http.StatusOK, envelope.New(true).WithData(map[string]any{
		"login": res,
	}))
}
