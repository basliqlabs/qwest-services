package httpserver

import (
	"net/http"

	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/labstack/echo/v4"
)

func (s *Server) healthCheck(c echo.Context) error {
	ctx := c.Request().Context()
	msg := contextutil.GetTranslation(ctx, "welcome", nil)
	return c.JSON(http.StatusOK, echo.Map{
		"message": msg,
	})
}
