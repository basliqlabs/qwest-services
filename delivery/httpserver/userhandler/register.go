package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/dto/userdto"
	"github.com/basliqlabs/qwest-services/pkg/echoutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

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
