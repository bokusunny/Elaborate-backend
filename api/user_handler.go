package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
)

// CreateUser called when request to '/users'
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	log.Printf("I got post request, json: " + string(body))

	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	newUser := entity.NewUser(user.Name, user.Email)
	database.DB.Create(&newUser)
	res := response{http.StatusOK, newUser}
	res.returnJSONToClient(w)
}
