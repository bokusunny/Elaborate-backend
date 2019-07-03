package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
)

type NewDirectoryResponse struct {
	Directory *entity.Directory `json:"newDirectory"`
}

type DirectoriesResponse struct {
	Directories []entity.Directory `json:"directories"`
}

// GET '/directories'
func FetchDirectoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.Header.Get("sub")
	log.Printf("[INFO] The uid for new directory: %v\n", userID)

	var directories []entity.Directory
	if err := database.DB.Where("user_id = ?", userID).Find(&directories).Error; err != nil {
		http.Error(w, "Error finding directories", http.StatusInternalServerError)
		return
	}

	res := DirectoriesResponse{directories}
	returnJSONToClient(w, res)
}

// POST '/directories'
func CreateDirectoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	log.Println("[INFO] Start creating new directory!")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] I got post request, json: %v\n", string(body))

	var directory entity.Directory
	if err := json.Unmarshal(body, &directory); err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusInternalServerError)
		return
	}

	userID := r.Header.Get("sub")
	log.Printf("[INFO] The uid for new directory: %v\n", userID)

	newDirectory := entity.NewDirectory(directory.Name, userID)
	if err := database.DB.Create(newDirectory).Error; err != nil {
		http.Error(w, "Error creating new directory", http.StatusInternalServerError)
		return
	}
	res := NewDirectoryResponse{newDirectory}
	returnJSONToClient(w, res)
}
