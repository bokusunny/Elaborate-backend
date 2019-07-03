package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Elaborate-backend/entity"
)

type Response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user"`
}

// interfaceで実装するのあんまよくないかも？
func returnJSONToClient(w http.ResponseWriter, r interface{}) {
	res, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] response : %s\n", string(res))

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
