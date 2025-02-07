package envelope

import (
	"errors"
	"net/http"

	"github.com/basliqlabs/qwest-services/pkg/contextutil"
	"github.com/basliqlabs/qwest-services/pkg/richerror"
	"github.com/basliqlabs/qwest-services/pkg/translation"
	"github.com/labstack/echo/v4"
)

func createResponse(code ErrorCode, err *richerror.RichError) *Response {
	return New(false).WithError(&ResponseError{
		Code:    code,
		Message: err.GetMessage(),
	})
}

func FromRichError(c echo.Context, err error) (int, *Response) {
	var re *richerror.RichError
	if errors.As(err, &re) {
		switch re.GetKind() {
		case richerror.KindInvalid:
			return http.StatusBadRequest, createResponse(ErrInvalidInput, re)
		case richerror.KindUnauthorized:
			return http.StatusUnauthorized, createResponse(ErrUnauthorized, re)
		case richerror.KindForbidden:
			return http.StatusForbidden, createResponse(ErrForbidden, re)
		case richerror.KindNotFound:
			return http.StatusNotFound, createResponse(ErrNotFound, re)
		// KindUUnexpected should not leak any data to the client, so it responds with a generic message.
		case richerror.KindUnexpected:
			lang := contextutil.GetLanguage(c.Request().Context())
			// TODO - aggregate this log with request log
			// logger.L().Named("error-handler").Error(re.GetMessage())
			return http.StatusInternalServerError, New(false).WithError(&ResponseError{
				Code:    ErrInternal,
				Message: translation.T(lang, "internal_server"),
			})
		default:
			return http.StatusInternalServerError, nil
		}
	}
	return http.StatusInternalServerError, nil
}
