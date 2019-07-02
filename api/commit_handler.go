package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Elaborate-backend/database"
	"github.com/Elaborate-backend/entity"
	"github.com/gorilla/mux"
)

type NewCommitResponse struct {
	Commit *entity.Commit `json:"newCommit"`
}

// POST '/directories/:directoryID/branches/:branchID/commits'
func CreateCommit(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// TODO: リクエストのbodyを読み込むの共通化できそう
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] I got post request, json: %v\n", string(body))

	var commit entity.Commit
	if err := json.Unmarshal(body, &commit); err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	branchID, err := strconv.Atoi(vars["branchID"])
	if err != nil {
		http.Error(w, "Cannot cast given branchID to int", http.StatusInternalServerError)
		return
	}
	log.Printf("[INFO] branchID: %v", branchID)

	newCommit := entity.NewCommit(commit.Name, commit.Body, branchID)
	if err := database.DB.Create(newCommit).Error; err != nil {
		http.Error(w, "Error creating new commit", http.StatusInternalServerError)
		return
	}

	res := NewCommitResponse{newCommit}
	returnJSONToClient(w, res)
}
