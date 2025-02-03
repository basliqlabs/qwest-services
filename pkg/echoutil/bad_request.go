package echoutil

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
)

func HandleBadRequest(c echo.Context) error {
	ctx := c.Request().Context()
	lang := contextutil.GetLanguage(ctx)

	return c.JSON(http.StatusBadRequest, envelope.New(false).WithError(&envelope.ResponseError{
		Code:    envelope.ErrBadRequest,
		Message: translation.T(lang, "bad_request"),
	}))
}
