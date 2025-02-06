package richerror

import (
	"errors"
)

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindUnexpected
	KindUnauthorized
	KindNotFound
)

type RichError struct {
	operation    string
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]any
}

func New(op string) *RichError {
	return &RichError{
		operation: op,
	}
}

func (r *RichError) WithMessage(msg string) *RichError {
	r.message = msg
	return r
}

func (r *RichError) WithError(err error) *RichError {
	r.wrappedError = err
	return r
}

func (r *RichError) WithKind(k Kind) *RichError {
	r.kind = k
	return r
}

func (r *RichError) WithMeta(meta map[string]any) *RichError {
	r.meta = meta
	return r
}

func (r *RichError) Error() string {
	if r.message == "" {
		if r.wrappedError != nil {
			return r.wrappedError.Error()
		}
		return "unspecified error"
	}
	return r.message
}

func (r *RichError) GetOperation() string {
	return r.operation
}

func (r *RichError) GetKind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	// RESEARCH: nestedRichError, ok := re.wrappedError.(RichError)
	var nestedRichError *RichError
	ok := errors.As(r.wrappedError, &nestedRichError)

	if !ok {
		return 0
	}

	return nestedRichError.GetKind()
}

func (r *RichError) GetMessage() string {
	if r.message != "" {
		return r.message
	}

	var nestedRichError *RichError
	ok := errors.As(r.wrappedError, &nestedRichError)

	if !ok {
		return r.wrappedError.Error()
	}

	return nestedRichError.GetMessage()
}

func (r *RichError) GetWrappedError() error {
	return r.wrappedError
}

func (r *RichError) GetMeta() map[string]any {
	return r.meta
}
