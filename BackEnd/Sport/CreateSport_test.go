package Sport

import (
	"BackEnd/DataBase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Setup an in-memory SQLite database for testing
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Auto-migrate schema
	db.AutoMigrate(&DataBase.Sport{})

	return db
}

// Test case for successfully creating a sport
func TestCreateSport(t *testing.T) {
	DataBase.DB = setupTestDB()

	sportRequest := map[string]interface{}{
		"Sport_name":        "Football",
		"Sport_Description": "A popular team sport",
	}

	body, _ := json.Marshal(sportRequest)
	req, err := http.NewRequest("POST", "/CreateSport", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSport)
	handler.ServeHTTP(recorder, req)

	t.Logf("Response Body: %s", recorder.Body.String())

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["message"] != "Sport record added successfully!!" {
		t.Errorf("unexpected response message: %v", response["message"])
	}

	var savedSport DataBase.Sport
	result := DataBase.DB.First(&savedSport, "sport_name = ?", "Football")

	if result.Error != nil {
		t.Errorf("Sport not found in database: %v", result.Error)
	}
}

// Test case for creating a duplicate sport
func TestCreateDuplicateSport(t *testing.T) {
	DataBase.DB = setupTestDB()

	// First request: should pass
	sportRequest := map[string]interface{}{
		"Sport_name":        "Basketball",
		"Sport_Description": "A popular team sport",
	}

	body, _ := json.Marshal(sportRequest)
	req, _ := http.NewRequest("POST", "/CreateSport", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateSport)
	handler.ServeHTTP(recorder, req)

	// Second request: should fail due to duplicate sport
	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/CreateSport", bytes.NewBuffer(body))
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}
