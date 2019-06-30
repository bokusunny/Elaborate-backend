package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
)

// POST '/users'
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	log.Printf("[INFO] I got post request, json: " + string(body))

	var user entity.User
	if err := json.Unmarshal(body, &user); err != nil {
		log.Fatal(err.Error())
	}

	newUser := entity.NewUser(user.Name, user.Email)
	database.DB.Create(&newUser)
	res := Response{http.StatusOK, newUser}
	returnJSONToClient(w, res)
}
