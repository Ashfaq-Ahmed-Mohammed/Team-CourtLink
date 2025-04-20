package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupListCourtsTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&DataBase.Sport{}, &DataBase.Court{})

	db.Create(&DataBase.Sport{
		Sport_name:        "Football",
		Sport_Description: "A popular team sport",
	})
	db.Create(&DataBase.Court{
		Court_Name:     "Court A",
		Court_Location: "Downtown",
		Court_Status:   1,
		Sport_id:       1,
	})

	return db
}

func TestListCourts(t *testing.T) {
	DataBase.DB = setupListCourtsTestDB()

	req, err := http.NewRequest("GET", "/ListCourts", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListCourts)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var courts []CourtData
	err = json.Unmarshal(recorder.Body.Bytes(), &courts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if len(courts) == 0 {
		t.Errorf("Expected courts to be listed, but got an empty response")
	}

	if courts[0].CourtName != "Court A" || courts[0].SportName != "Football" {
		t.Errorf("unexpected court data: %+v", courts)
	}
}

func TestListMultipleCourts(t *testing.T) {
	db := setupListCourtsTestDB()
	DataBase.DB = db

	db.Exec("DELETE FROM Court")

	courts := []DataBase.Court{
		{Court_Name: "Test Court A", Court_Location: "Rietzz", Court_Status: 1, Sport_id: 1},
		{Court_Name: "Test Court B", Court_Location: "Southwest", Court_Status: 1, Sport_id: 1},
	}
	if err := db.Create(&courts).Error; err != nil {
		t.Fatalf("failed to insert test courts: %v", err)
	}

	req, err := http.NewRequest("GET", "/ListCourts", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListCourts)
	handler.ServeHTTP(recorder, req)

	resp := recorder.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, resp.StatusCode)
	}

	var courtsResponse []CourtData
	if err := json.NewDecoder(resp.Body).Decode(&courtsResponse); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(courtsResponse) != 2 {
		t.Errorf("expected 2 courts, got %d", len(courtsResponse))
	}

	expectedNames := []string{"Test Court A", "Test Court B"}
	for i, court := range courtsResponse {
		if court.CourtName != expectedNames[i] || court.SportName != "Football" {
			t.Errorf("unexpected court data at index %d: %+v", i, court)
		}
	}
}

func TestListCourtsEmpty(t *testing.T) {
	DataBase.DB = setupListCourtsTestDB()

	DataBase.DB.Exec("DELETE FROM Court")

	req, err := http.NewRequest("GET", "/ListCourts", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListCourts)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var courts []CourtData
	err = json.Unmarshal(recorder.Body.Bytes(), &courts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if len(courts) != 0 {
		t.Errorf("Expected 0 courts to be listed, but got %d", len(courts))
	}
}
