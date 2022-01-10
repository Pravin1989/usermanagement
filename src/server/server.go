package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"ecommerce/user/usermanagement/src/config"

	"github.com/gorilla/mux"
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
