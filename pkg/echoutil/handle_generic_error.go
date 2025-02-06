package echoutil

import (
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

func HandleGenericError(c echo.Context, err error) error {
	return c.JSON(envelope.FromRichError(c, err))
}
