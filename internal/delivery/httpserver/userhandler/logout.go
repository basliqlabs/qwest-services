package userhandler

import (
	"net/http"

	"github.com/basliqlabs/qwest-services/pkg/echoutil"
	"github.com/basliqlabs/qwest-services/pkg/envelope"
	"github.com/labstack/echo/v4"
)

// logout lets a user logout
//
//	@Summary		User logout
//	@Description	Logout a user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	envelope.OpenAPIResponseSuccess{data=tokendto.LogoutResponse}
//	@Failure		400		{object}	envelope.OpenAPIResponseError
//	@Failure		422		{object}	envelope.OpenAPIResponseError
//	@Router			/users/logout [post]
func (h Handler) logout(c echo.Context) error {
	ctx := c.Request().Context()
	email, ok := c.Get("email").(string)
	if !ok {
		return echoutil.HandleBadRequest(c)
	}
	err := h.tokenService.Logout(ctx, email)
	if err != nil {
		return echoutil.HandleGenericError(c, err)
	}
	return c.JSON(http.StatusOK, envelope.New(true))
}
