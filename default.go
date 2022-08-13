package errors

const (
	AlreadyExists  ErrorCode = 1
	NotFound       ErrorCode = 2
	Internal       ErrorCode = 3
	InvalidRequest ErrorCode = 4
)

var Default = map[ErrorCode]string{
	AlreadyExists:  "registry already exists",
	NotFound:       "not found",
	InvalidRequest: "invalid request",
	Internal:       "unexpected error has occurred",
}
