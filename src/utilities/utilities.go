package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func EncodeResponse(r *http.Request, w http.ResponseWriter, resObj interface{}) {
	err := json.NewEncoder(w).Encode(&resObj)
	if err != nil {
		fmt.Println("Error Occured : ", err)
	}
}

func GetURLParam(r *http.Request, key string) string {
	params := mux.Vars(r)
	return params[key]
}

func DecodeRequest(r *http.Request, object interface{}) error {
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&object)
}

func GenerateUUID() uuid.UUID {
	return uuid.New()
}
