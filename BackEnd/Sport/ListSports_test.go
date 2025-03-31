package Sport

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListSports(t *testing.T) {
	DataBase.DB = setupTestDB()

	sport := DataBase.Sport{Sport_name: "Football"}
	DataBase.DB.Create(&sport)

	req, err := http.NewRequest("GET", "/ListSports", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListSports)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var response []string
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	expectedSport := "Football"
	if len(response) == 0 || response[0] != expectedSport {
		t.Errorf("expected %v, got %v", expectedSport, response)
	}
}
