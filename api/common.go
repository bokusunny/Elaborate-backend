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

func (r *Response) returnJSONToClient(w http.ResponseWriter) {
	res, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("response : %s\n", string(res))

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
