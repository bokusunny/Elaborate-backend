package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
)

type DirectoryResponse struct {
	Status    int               `json:"status"`
	Directory *entity.Directory `json:"newDirectory"`
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
	}

	log.Printf("[INFO] I got post request, json: %v\n", string(body))

	var directory entity.Directory
	if err := json.Unmarshal(body, &directory); err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusInternalServerError)
	}

	userID := r.Header.Get("sub")
	log.Printf("[INFO] The uid for new directory: %v\n", userID)

	newDirectory := entity.NewDirectory(directory.Name, userID)
	database.DB.Create(&newDirectory)
	res := DirectoryResponse{http.StatusOK, newDirectory}
	returnJSONToClient(w, res)
}
