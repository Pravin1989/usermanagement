package utilities

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeResponse(r *http.Request, w http.ResponseWriter, resObj interface{}) {
	err := json.NewEncoder(w).Encode(&resObj)
	if err != nil {
		fmt.Println("Error Occured : ", err)
	}
}
