package middleware

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/basliqlabs/qwest-services-auth/pkg/envelope"
	"github.com/basliqlabs/qwest-services-auth/pkg/logger"
	"github.com/basliqlabs/qwest-services-auth/translation"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Recovery(t *translation.Translator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// TODO: this is a temporary solution to pass the translator to the recovery middleware
			ctx := contextutil.WithTranslator(c.Request().Context(), t)
			lang := c.Request().Header.Get("Accept-Language")
			ctx = contextutil.WithLanguage(ctx, lang)
			c.SetRequest(c.Request().WithContext(ctx))

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
						Message: contextutil.GetTranslation(ctx, "internal_server", nil),
					}))
				}
			}()

			return next(c)
		}
	}
}
