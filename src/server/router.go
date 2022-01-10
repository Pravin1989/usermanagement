package server

import (
	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/entity/models"
	"ecommerce/user/usermanagement/src/handlers"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	HEADER_CONTENT_TYPE_VALUE = "application/json; charset=utf-8"
	HEADER_CONTENT_TYPE_KEY   = "Content-Type"
)

var (
	handleGetUserDetails  = handlers.HandleGetUserDetails
	handleSaveUserDetails = handlers.HandleSaveUserDetails
)

func registerServiceRoutes(r *mux.Router) {
	r.HandleFunc("/user/fetch/{userId}", handleGetUserDetails).Methods(http.MethodGet)
	r.HandleFunc("/user/add", handleSaveUserDetails).Methods(http.MethodPost)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}

//This method retrieves the transcationID from request,
//If present adds it to context if not then it generates a new transactionID before adding to context.
//It also sets to the transcationID (either retrieved or generated) into response header
func transactionIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		transactionID := r.Header.Get(models.TransactionIDHeaderKey)
		ctx = config.SetTransactionIDInContext(ctx, transactionID)
		w.Header().Add(models.CorrelationIDHeaderKey, config.TransactionIDFromContext(ctx))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func registerMiddleWares(router *mux.Router) {
	router.Use(headerMiddleWare, transactionIDMiddleware)
}

// Adds response headers before calling the the controller
func headerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(HEADER_CONTENT_TYPE_KEY, HEADER_CONTENT_TYPE_VALUE)
		next.ServeHTTP(w, r)
	})
}
