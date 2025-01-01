package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/basliqlabs/qwest-services-auth/pkg/envelope"
	"github.com/basliqlabs/qwest-services-auth/pkg/translation"
	"github.com/basliqlabs/qwest-services-auth/service/authservice"
	"github.com/labstack/echo/v4"
)

// userLogin lets a user get an authorization token
//
//	@Summary		User login
//	@Description	Authenticate a user with username and password
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		dto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=dto.LoginResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/login [post]
func (h Handler) userLogin(c echo.Context) error {
	ctx := c.Request().Context()
	lang := contextutil.GetLanguage(ctx)
	req := new(dto.LoginRequest)

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

	// RESEARCH: pointer vs concrete structs
	authSvc := authservice.New("")
	res, err := authSvc.Login(ctx, req)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, envelope.FromRichError(err))
	}

	return c.JSON(http.StatusOK, envelope.New(true).WithData(map[string]any{
		"login": res,
	}))
}
