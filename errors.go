package errors

import "fmt"

const (
	AlreadyExists  ErrorCode = 1
	NotFound       ErrorCode = 2
	Internal       ErrorCode = 3
	InvalidRequest ErrorCode = 4
)

var (
	errors = map[ErrorCode]string{
		AlreadyExists:  "registry already exists",
		NotFound:       "not found",
		InvalidRequest: "invalid request",
		Internal:       "unexpected error has occurred",
	}
)

type ErrorCode int

type Error struct {
	Code    ErrorCode
	Message string
	Details interface{}
}

func Setup(code ErrorCode, msg string) error {
	if _, ok := errors[code]; ok {
		panic(fmt.Sprintf("error already exists %d", code))
	}

	errors[code] = msg
	return nil
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d - err: %s", e.Code, e.Message)
}

func (e *Error) Is(code ErrorCode) bool {
	return e.Code == code
}

func Get(err error) (*Error, bool) {
	if libErr, ok := err.(*Error); ok {
		return libErr, true
	}
	return nil, false
}

func Code(code ErrorCode, msg ...string) *Error {
	var errMsg string
	var ok bool
	if len(msg) > 0 {
		errMsg = msg[0]
	} else {
		errMsg, ok = errors[code]
		if !ok {
			panic(fmt.Sprintf("error not found with code %d", code))
		}
	}

	return &Error{
		Code:    code,
		Message: errMsg,
	}
}

func Detailed(code ErrorCode, details interface{}, msg ...string) *Error {
	err := Code(code, msg...)
	err.Details = details
	return err
}
