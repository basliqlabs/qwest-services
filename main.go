package main

import (
	"github.com/basliqlabs/qwest-services-auth/dto"
	"github.com/basliqlabs/qwest-services-auth/service/authservice"
	"github.com/basliqlabs/qwest-services-auth/validator/authvalidator"
	"github.com/labstack/echo/v4"
	"net/http"
)

// TODO: config
// TODO: repository and migration
// TODO: rich error
// TODO: hot reload
// TODO: logging
// TODO: envelope

func main() {
	e := echo.New()
	e.POST("/login", func(c echo.Context) error {
		req := new(dto.LoginRequest)
		if err := c.Bind(req); err != nil {
			// TODO: bad request handler
			return c.String(http.StatusBadRequest, "bad request")
		}

		validationErrors, err := authvalidator.Login(req)
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
	})

	e.Logger.Fatal(e.Start(":1323"))
}
