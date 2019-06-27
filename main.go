package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Elaborate-backend/api"
	"github.com/Elaborate-backend/database"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db := database.DB
	defer db.Close()

	// TODO: originは環境によって場合分け
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})
	allowedHeaders := handlers.AllowedHeaders([]string{"Authorization", "Content-Type"})

	r := mux.NewRouter()
	r.HandleFunc("/", api.CreateUser) // TODO: "/users"に変更 && POSTリクエストに限定

	log.Fatal(http.ListenAndServe(":"+os.Getenv("port"), handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(r)))
}
