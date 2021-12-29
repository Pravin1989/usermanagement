package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/handlers"

	"github.com/gorilla/mux"
)

var (
	handleGetUserDetails = handlers.HandleGetUserDetails
)

const (
	port = 8080
)

func Start(context context.Context) {
	transactionID := config.TransactionIDFromContext(context)
	router := loadRoutes()
	log.Printf("Starting REST Server, TransactionID : %v", transactionID)
	log.Printf("REST server listening on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), removeTrailingSlash(router))
	fmt.Println("Server Crashed", err)
}

func loadRoutes() *mux.Router {
	serviceRouter := mux.NewRouter().PathPrefix("/usermanagement/api").Subrouter()
	registerMiddleWares(serviceRouter)
	registerServiceRoutes(serviceRouter)
	return serviceRouter
}
func registerServiceRoutes(r *mux.Router) {
	r.HandleFunc("/user/fetch/{userId}", handleGetUserDetails).Methods(http.MethodGet)
}

func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
