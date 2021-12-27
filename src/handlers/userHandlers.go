package handlers

import (
	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/utilities"
	"fmt"
	"net/http"
)

func HandleGetUserDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transcationID := config.TransactionIDFromContext(ctx)
	fmt.Println("Started processing of GetUserDetailsByID", transcationID)
	utilities.EncodeResponse(r, w, "Hello Pravin Budage")
}
