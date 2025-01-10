package envelope

type ErrorCode string

const (
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrForbidden    ErrorCode = "FORBIDDEN"
	ErrInvalidToken ErrorCode = "INVALID_TOKEN"
	ErrTokenExpired ErrorCode = "TOKEN_EXPIRED"

	ErrValidation    ErrorCode = "VALIDATION_ERROR"
	ErrInvalidInput  ErrorCode = "INVALID_INPUT"
	ErrMalformedJSON ErrorCode = "MALFORMED_JSON"

	ErrNotFound      ErrorCode = "NOT_FOUND"
	ErrAlreadyExists ErrorCode = "ALREADY_EXISTS"
	ErrConflict      ErrorCode = "CONFLICT"
	ErrBadRequest    ErrorCode = "BAD_REQUEST"

	ErrDatabase   ErrorCode = "DATABASE_ERROR"
	ErrConnection ErrorCode = "CONNECTION_ERROR"

	ErrInternal    ErrorCode = "INTERNAL_ERROR"
	ErrServiceDown ErrorCode = "SERVICE_DOWN"
	ErrTimeout     ErrorCode = "TIMEOUT"

	ErrInvalidOperation ErrorCode = "INVALID_OPERATION"
	ErrLimitExceeded    ErrorCode = "LIMIT_EXCEEDED"
)
