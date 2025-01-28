package httpserver

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
)

// HealthCheck godoc
//
//	@Summary		Health check
//	@Description	Check if the API is running
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	map[string]string
//	@Router			/healthcheck [get]
func (s *Server) healthCheck(c echo.Context) error {
	lang := contextutil.GetLanguage(c.Request().Context())

	return c.JSON(http.StatusOK, echo.Map{
		"message": translation.T(lang, "welcome", nil),
	})
}
