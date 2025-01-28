package middleware

import (
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/labstack/echo/v4"
)

func TranslatorMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			lang := req.Header.Get("Accept-Language")
			ctx := contextutil.WithLanguage(req.Context(), lang)
			c.SetRequest(req.WithContext(ctx))
			return next(c)
		}
	}
}
