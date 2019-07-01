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

type BranchResponse struct {
	Branch entity.Branch `json:"branch"`
}

// GET '/directories/:directoryID/branches'
func FetchOpenBranchesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	directoryID, err := strconv.Atoi(vars["directoryID"])
	if err != nil {
		http.Error(w, "Cannot cast given directoryID to int", http.StatusInternalServerError)
		return
	}

	var directory entity.Directory
	if database.DB.First(&directory, directoryID).RecordNotFound() {
		http.Error(w, "Directory not found with given id", http.StatusInternalServerError)
		return
	}

	var branches []entity.Branch
	database.DB.Model(&directory).Related(&branches).Where("state = ?", "open").Find(&branches)

	res := BranchesResponse{branches}
	returnJSONToClient(w, res)
}

// POST '/directories/:directoryID/branches'
func CreateBranchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	log.Printf("[INFO] I got post request, json: %v\n", string(body))

	var branch entity.Branch
	if err := json.Unmarshal(body, &branch); err != nil {
		http.Error(w, "Error unmarshal request body", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	directoryID, err := strconv.Atoi(vars["directoryID"])
	log.Printf("[INFO] directoryID: %v", directoryID)
	if err != nil {
		http.Error(w, "Cannot cast given directoryID to int", http.StatusInternalServerError)
		return
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

// GET '/directories/:directoryID/branches/:branchID
func FetchBranchWithIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	vars := mux.Vars(r)
	directoryID, err := strconv.Atoi(vars["directoryID"])
	if err != nil {
		http.Error(w, "Cannot cast given directoryID to int", http.StatusInternalServerError)
		return
	}
	branchID, err := strconv.Atoi(vars["branchID"])
	if err != nil {
		http.Error(w, "Cannot cast given branchID to int", http.StatusInternalServerError)
		return
	}
	log.Printf("[INFO] directoryID: %v, branchID: %v", directoryID, branchID)

	var directory entity.Directory
	if database.DB.First(&directory, directoryID).RecordNotFound() {
		http.Error(w, "Directory not found with given id", http.StatusInternalServerError)
		return
	}

	var branches []entity.Branch
	var branch entity.Branch
	if database.DB.Model(&directory).Related(&branches).First(&branch, branchID).RecordNotFound() {
		http.Error(w, "Branch not found with given id", http.StatusNotFound)
		return
	}

	res := BranchResponse{branch}
	returnJSONToClient(w, res)
}
