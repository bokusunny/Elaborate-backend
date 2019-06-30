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
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateUserHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("[ERROR] handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var res api.Response
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatal(err.Error())
	}

	// 標準パッケージではtimeをモックできないのでresponseのうちcreatedAt, updatedAtは検証しない
	// TODO: User structのプロパティが増えるとここもいちいち書き換えないといけないのでなんとかする
	if res.User.Name != testData["name"] {
		t.Errorf("[ERROR] returned user has unexpected name: got %v want %v",
			res.User.Name, testData["name"])
	}
	if res.User.Email != testData["email"] {
		t.Errorf("[ERROR] returned user has unexpected name: got %v want %v",
			res.User.Email, testData["email"])
	}
}
