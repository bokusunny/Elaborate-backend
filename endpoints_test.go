package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Elaborate-backend/api"
)

func TestCreateDirectoryHandler(t *testing.T) {
	testData := map[string]string{"name": "testDir"}
	testBody, _ := json.Marshal(testData)

	req, err := http.NewRequest("POST", "/directories", bytes.NewBuffer(testBody))
	if err != nil {
		t.Fatal(err.Error())
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.CreateDirectoryHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("[ERROR] handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var res api.DirectoryResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &res); err != nil {
		t.Fatal(err.Error())
	}

	// 標準パッケージではtimeをモックできないのでresponseのうちcreatedAt, updatedAtは検証しない
	// TODO: Directory structのプロパティが増えるとここもいちいち書き換えないといけないのでなんとかする
	if res.Directory.Name != testData["name"] {
		t.Errorf("[ERROR] returned directory has unexpected name: got %v want %v",
			res.Directory.Name, testData["name"])
	}
}
