package errors

import "fmt"

var errors = map[ErrorCode]string{}

type ErrorCode int

type Error struct {
	code    ErrorCode
	message string
	details interface{}
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
	return fmt.Sprintf("code: %d - err: %s", e.code, e.message)
}

func (e *Error) Code() ErrorCode {
	return e.code
}

func (e *Error) Details() interface{} {
	return e.details
}

func (e *Error) Is(code ErrorCode) bool {
	return e.code == code
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
		code:    code,
		message: errMsg,
	}
}

func Detailed(code ErrorCode, details interface{}, msg ...string) *Error {
	err := Code(code, msg...)
	err.details = details
	return err
}
