package api

import (
	"encoding/json"
	"log"
	"net/http"
)

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
