package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type response struct {
	Status int          `json:"status"`
	User   *entity.User `json:"user"`
}

func (r *response) returnJSONToClient(w http.ResponseWriter) {
	res, err := json.Marshal(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("response : %s\n", string(res))

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	db := database.DB
	defer db.Close()

	createUser := func(w http.ResponseWriter, r *http.Request) {
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
		db.Create(&newUser)
		res := response{http.StatusOK, newUser}
		res.returnJSONToClient(w)
	}

	// TODO: originは環境によって場合分け
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

	r := mux.NewRouter()
	r.HandleFunc("/", createUser)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
