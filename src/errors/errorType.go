package errors

import "net/http"

type ErrorType struct {
	typeString string
	statusCode statusCodeWithPriority
}

type statusCodeWithPriority struct {
	code     int
	priority int
}

var (
	UnexpectedError = ErrorType{typeString: "UnexpectedError", statusCode: internalServerError}
)
var (
	internalServerError = statusCodeWithPriority{code: http.StatusInternalServerError, priority: 1}
)

func (e ErrorType) Type() string {
	return e.typeString
}
