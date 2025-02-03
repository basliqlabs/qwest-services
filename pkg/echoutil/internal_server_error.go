package echoutil

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

func HandleInternalServerError(c echo.Context, err error) error {
	return c.JSON(http.StatusUnauthorized, envelope.FromRichError(err))
}
