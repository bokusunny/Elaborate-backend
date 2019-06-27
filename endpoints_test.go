package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Elaborate-backend/api"
)

func TestCreateUserHandler(t *testing.T) {
	testData := map[string]string{"name": "ninjawanko", "email": "ninjawanko@example.com"}
	testBody, _ := json.Marshal(testData)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateUserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var res api.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatal("failed to unmarshal response.")
	}

	// 標準パッケージではtimeをモックできないのでresponseのうちcreatedAt, updatedAtは検証しない
	if res.User.Name != testData["name"] {
		t.Errorf("returned user has unexpected name: got %v want %v",
			res.User.Name, testData["name"])
	}
	if res.User.Email != testData["email"] {
		t.Errorf("returned user has unexpected name: got %v want %v",
			res.User.Email, testData["email"])
	}
}