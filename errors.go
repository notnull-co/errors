package errors

import "fmt"

var errors = map[ErrorCode]string{}

type ErrorCode int

type Error struct {
	Code    ErrorCode   `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func SetupMulti(codes map[ErrorCode]string) {
	for code, msg := range codes {
		Setup(code, msg)
	}
}

func Setup(code ErrorCode, msg string) {
	if _, ok := errors[code]; ok {
		panic(fmt.Sprintf("error already exists %d", code))
	}
	errors[code] = msg
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
