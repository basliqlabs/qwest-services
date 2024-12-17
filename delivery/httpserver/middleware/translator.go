package middleware

import (
	"github.com/basliqlabs/qwest-services-auth/pkg/contextutil"
	"github.com/basliqlabs/qwest-services-auth/translation"
	"github.com/labstack/echo/v4"
)

func TranslatorMiddleware(t *translation.Translator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := contextutil.WithTranslator(c.Request().Context(), t)
			lang := c.Request().Header.Get("Accept-Language")
			ctx = contextutil.WithLanguage(ctx, lang)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
