package userhandler

import (
	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/service/authservice"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) userLogin(c echo.Context) error {
	req := new(dto.LoginRequest)
	if err := c.Bind(req); err != nil {
		// TODO: bad request handler
		return c.String(http.StatusBadRequest, "bad request")
	}

	// TODO: temporarily using nil as context
	validationErrors, err := h.validator.Login(nil, req)

	if err != nil {
		// TODO: implement field errors
		return c.JSON(http.StatusUnprocessableEntity, validationErrors)
	}

	// RESEARCH: pointer vs concrete structs
	// TODO: fix interface{}
	authSvc := authservice.New("")
	res, err := authSvc.Login(req)
	if err != nil {
		return c.String(http.StatusInternalServerError, "server error")
	}

	return c.JSON(http.StatusOK, res)
}
