package handlers

import (
	"ecommerce/user/usermanagement/src/common"
	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/entity/models"
	"ecommerce/user/usermanagement/src/service"
	"ecommerce/user/usermanagement/src/utilities"
	"fmt"
	"net/http"
)

var (
	getURLParam        = utilities.GetURLParam
	decodeRequest      = utilities.DecodeRequest
	encodeResponse     = utilities.EncodeResponse
	getUserDetailsById = service.GetUserDetailsById
	saveUserDetail     = service.SaveUserDetails
)

func HandleGetUserDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transcationID := config.TransactionIDFromContext(ctx)
	fmt.Println("Started processing of GetUserDetailsByID", transcationID)
	userId := getURLParam(r, common.UserIdStr)
	response, err := getUserDetailsById(models.UserId{ID: userId})
	if err != nil {
		fmt.Println("Failed to fetch data from DB", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	utilities.EncodeResponse(r, w, response)
	fmt.Println("Completed processing of GetUserDetailsByID", transcationID)
}

func HandleSaveUserDetails(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transcationID := config.TransactionIDFromContext(ctx)
	fmt.Println("Started processing of HandleSaveUserDetails", transcationID)
	var userRequest models.UserRequest
	decodeError := decodeRequest(r, &userRequest)
	if decodeError != nil {
		fmt.Printf("Failed to decode %v", decodeError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user := userRequest.ToUser()
	userSaved, saveError := saveUserDetail(user)
	if saveError != nil {
		fmt.Printf("Failed to decode %v", saveError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	encodeResponse(r, w, &userSaved)
	fmt.Println("Completed processing of HandleSaveUserDetails", transcationID)
}
