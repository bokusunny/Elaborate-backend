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

type NewBranchResponse struct {
	Branch *entity.Branch `json:"newBranch"`
}

type BranchesResponse struct {
	Branches []entity.Branch `json:"branches"`
}

// GET '/directories/:directoryID/branches'
func FetchOpenBranchesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	vars := mux.Vars(r)
	directoryID, err := strconv.Atoi(vars["directoryID"])
	if err != nil {
		http.Error(w, "Invalid directory id", http.StatusInternalServerError)
	}

	var directory entity.Directory
	database.DB.First(&directory, directoryID)

	var branches []entity.Branch
	database.DB.Model(&directory).Related(&branches).Where("state = ?", "open").Find(&branches)

	res := BranchesResponse{branches}
	returnJSONToClient(w, res)
}

// POST '/directories/:directoryID/branches'
func CreateBranchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
	}

	log.Printf("[INFO] I got post request, json: %v\n", string(body))

	var branch entity.Branch
	if err := json.Unmarshal(body, &branch); err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusInternalServerError)
	}

	vars := mux.Vars(r)
	directoryID, err := strconv.Atoi(vars["directoryID"])
	if err != nil {
		http.Error(w, "Invalid directory id", http.StatusInternalServerError)
	}

	newBranch := entity.NewBranch(
		branch.Name,
		directoryID,
		branch.BaseBranchID,
		branch.BaseBranchName,
		branch.Body,
		branch.State,
	)
	database.DB.Create(newBranch)
	res := NewBranchResponse{newBranch}
	returnJSONToClient(w, res)
}
