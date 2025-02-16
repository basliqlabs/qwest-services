package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/echoutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

// register lets a user register
//
//	@Summary		User register
//	@Description	Register with an email and a password
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		userdto.RegisterRequest	true	"Register credentials"
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=userdto.RegisterResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/register [post]
func (h Handler) register(c echo.Context) error {
	ctx := c.Request().Context()
	req := new(userdto.RegisterRequest)
	if err := c.Bind(&req); err != nil {
		return echoutil.HandleBadRequest(c)
	}
	validationErrors, err := h.validator.Register(ctx, req)
	if err != nil {
		return echoutil.HandleUnprocessableContent(c, validationErrors)
	}
	res, err := h.service.Register(ctx, req)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true).WithData(res))
}
