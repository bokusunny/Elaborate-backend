package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// ResponseData should have tag?
type ResponseData struct {
	StatusCode int
	Data       string
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Elaborate-backend!</h1>")
}

func omikujiHandler(w http.ResponseWriter, r *http.Request) {
	oracles := []string{"大吉", "中吉", "小吉", "末吉", "吉", "凶", "大凶"}
	response := ResponseData{http.StatusOK, oracles[rand.Intn(7)]}

	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/omikuji", omikujiHandler)
	http.ListenAndServe(":3000", nil)
}
