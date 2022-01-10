package errors

import (
	"context"
	"ecommerce/user/usermanagement/src/entity/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleError(ctx context.Context, req *http.Request, res http.ResponseWriter, err error) {
	responseError, statusCode := createResponseError(req, err)
	responseErrors := make([]models.Error, 1)
	responseErrors[0] = responseError
	sendError(res, models.ErrorResponse{Errors: responseErrors}, statusCode.code)
}
func createResponseError(req *http.Request, err error) (models.Error, statusCodeWithPriority) {
	serviceError, ok := err.(userServiceError)
	if !ok {
		serviceError = newErrorObject(UnhandledError, UnexpectedError)
	}
	message := serviceError.translateErrorMessage()
	responseError := models.Error{Code: string(serviceError.getErrorCode()), Message: message}
	return responseError, serviceError.errorType.statusCode
}
func sendError(w http.ResponseWriter, errorMessge models.ErrorResponse, statusCode int) {
	data, err := json.Marshal(errorMessge)
	if err != nil {
		http.Error(w, UnexpectedError.Type(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	fmt.Fprintln(w, string(data))
}
