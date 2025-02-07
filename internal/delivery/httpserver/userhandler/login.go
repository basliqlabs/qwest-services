package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/echoutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

// login lets a user get an authorization token
//
//	@Summary		User login
//	@Description	Authenticate a user with username/phone/email and password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		userdto.LoginRequest	true	"Login credentials"
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=userdto.LoginResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/login [post]
func (h Handler) login(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(userdto.LoginRequest)
	if err := c.Bind(req); err != nil {
		return echoutil.HandleBadRequest(c)
	}
	validationErrors, err := h.validator.Login(ctx, req)
	if err != nil {
		return echoutil.HandleUnprocessableContent(c, validationErrors)
	}
	res, err := h.service.Login(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
