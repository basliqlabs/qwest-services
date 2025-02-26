package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/basliqlabs/qwest-services/internal/dto/userdto"
	"github.com/basliqlabs/qwest-services/internal/service/tokenservice"
	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/basliqlabs/qwest-services/pkg/jwtutil"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
)

type AuthMiddleware struct {
	tokenRepo tokenservice.Repository
}

func NewAuthMiddleware(tokenRepo tokenservice.Repository) AuthMiddleware {
	return AuthMiddleware{
		tokenRepo: tokenRepo,
	}
}

// StrictAuth validates tokens in the Authorization header and adds the user to context
func (m *AuthMiddleware) StrictAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			lang := contextutil.GetLanguage(c.Request().Context())
			if authHeader == "" {
				return c.JSON(
					http.StatusUnauthorized,
					envelope.New(false).
						WithError(
							&envelope.ResponseError{
								Code:    envelope.ErrUnauthorized,
								Message: translation.T(lang, "unauthorized"),
							},
						),
				)
			}

			// Extract token from "Bearer <token>"
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(
					http.StatusUnauthorized,
					envelope.New(false).
						WithError(
							&envelope.ResponseError{
								Code:    envelope.ErrUnauthorized,
								Message: translation.T(lang, "invalid_access_token"),
							},
						),
				)
			}
			tokenString := parts[1]

			// Validate token
			claims, err := jwtutil.Decode(tokenString)
			if err != nil {
				return c.JSON(
					http.StatusUnauthorized,
					envelope.New(false).
						WithError(
							&envelope.ResponseError{
								Code:    envelope.ErrUnauthorized,
								Message: translation.T(lang, "invalid_token_claims"),
							},
						),
				)
			}

			// Check if token is revoked
			refreshToken, err := m.tokenRepo.GetByToken(c.Request().Context(), tokenString)
			if err == nil && refreshToken.Revoked {
				return c.JSON(
					http.StatusUnauthorized,
					envelope.New(false).
						WithError(
							&envelope.ResponseError{
								Code:    envelope.ErrUnauthorized,
								Message: translation.T(lang, "token_has_been_revoked"),
							},
						),
				)
			}

			// Add user info to context
			if username, ok := claims["username"].(string); ok {
				c.Set("username", username)
			}
			if email, ok := claims["email"].(string); ok {
				c.Set("email", email)
			}

			return next(c)
		}
	}
}

// OptionalAuth checks for a token but doesn't require it
// This is useful for login endpoints where you want to skip
// token generation if a valid token exists
func (m *AuthMiddleware) OptionalAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Extract token from header if it exists
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader != "" {
				parts := strings.Split(authHeader, " ")
				if len(parts) == 2 && parts[0] == "Bearer" {
					tokenString := parts[1]
					claims, err := jwtutil.Decode(tokenString)
					if err == nil {
						if exp, ok := claims["exp"].(float64); ok {
							if time.Unix(int64(exp), 0).After(time.Now()) {
								refreshToken, err := m.tokenRepo.GetByUserEmail(c.Request().Context(), claims["email"].(string))
								if err == nil && !refreshToken.Revoked {
									return c.JSON(
										http.StatusOK,
										envelope.New(true).
											WithData(&userdto.LoginResponse{
												AccessToken:  tokenString,
												RefreshToken: refreshToken.Token,
											}))
								}
							}
						}
					}
				}
			}

			return next(c)
		}
	}
}
