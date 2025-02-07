package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/basliqlabs/qwest-services/pkg/logger"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Recovery() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lang := c.Request().Header.Get("Accept-Language")

			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}

					stack := make([]byte, 4<<10)
					length := runtime.Stack(stack, false)

					fields := []zap.Field{
						zap.String("panic", err.Error()),
						zap.String("stack", string(stack[:length])),
						zap.String("method", c.Request().Method),
						zap.String("path", c.Request().URL.Path),
						zap.String("remote_ip", c.RealIP()),
					}

					logger.L().Named("panic").Error("recovered from panic", fields...)

					c.JSON(http.StatusInternalServerError, envelope.New(false).WithError(&envelope.ResponseError{
						Code:    envelope.ErrInternal,
						Message: translation.T(lang, "internal_server"),
					}))
				}
			}()

			return next(c)
		}
	}
}
