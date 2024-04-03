package internal

import (
	"fmt"
)

type ErrorCode uint

type Error struct {
	code    ErrorCode
	message string
	origin  error
	code    string
}

const (
	ErrorCodeUnknown ErrorCode = iota
	ErrorCodeNotFound
	ErrorCodeInvalidArgument
)

func WrapError(origin error, code ErrorCode, format string, a ...interface{}) error {
	return &Error{
		code:    code,
		message: fmt.Sprintf(format, a...),
		origin:  origin,
	}
}

func NewError(code ErrorCode, format string, a ...interface{}) error {
	return WrapError(nil, code, format, a...)
}

func (e *Error) Error() string {
	if e.origin != nil {
		return fmt.Sprintf("%s: %v", e.message, e.origin)
	}

	return e.message
}

func (e *Error) UnWrap() error {
	return e.origin
}

func (e *Error) Code() ErrorCode {
	return e.code
}
