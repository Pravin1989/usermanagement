package errors

import (
	"ecommerce/user/usermanagement/src/translation"
)

type userServiceError struct {
	errorType ErrorType
	errorCode ErrorCode
	message   string
	params    []interface{}
}

func (e userServiceError) Error() string {
	return e.message
}
func (e userServiceError) getErrorCode() ErrorCode {
	return e.errorCode
}
func (e userServiceError) getErrorType() ErrorType {
	return e.errorType
}
func newErrorObject(errorCode ErrorCode, errorType ErrorType, params ...interface{}) userServiceError {
	return userServiceError{
		errorCode: errorCode,
		errorType: errorType,
		params:    params,
		message:   string(errorCode),
	}
}

func (e userServiceError) translateErrorMessage() string {
	translatedMessage, err := translation.TraanslateToDefaultlanguage(e.message, e.params...)
	if err != nil || translatedMessage == "" {
		translatedMessage = e.message
	}
	return translatedMessage
}
