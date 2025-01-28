package envelope

import (
	"errors"

	"github.com/basliqlabs/qwest-services/pkg/richerror"
)

func createResponse(code ErrorCode, err *richerror.RichError) *Response {
	return New(false).WithError(&ResponseError{
		Code:    code,
		Message: err.GetMessage(),
	})
}

func FromRichError(err error) *Response {
	var re *richerror.RichError
	if errors.As(err, &re) {
		switch re.GetKind() {
		case richerror.KindInvalid:
			return createResponse(ErrInvalidInput, re)
		case richerror.KindForbidden:
			return createResponse(ErrForbidden, re)
		case richerror.KindUnexpected:
			return createResponse(ErrInternal, re)
		case richerror.KindUnauthorized:
			return createResponse(ErrUnauthorized, re)
		case richerror.KindNotFound:
			return createResponse(ErrNotFound, re)
		default:
			// TODO: this should not happen
			return nil
		}
	}
	return nil
}

func BadRequest() *Response {
	return New(false).WithError(&ResponseError{
		Code:    ErrBadRequest,
		Message: "",
	})
}
