package Court

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

func setupTestDBForCreateCourt() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&DataBase.Sport{}, &DataBase.Court{}, &DataBase.Court_TimeSlots{})

	db.Create(&DataBase.Sport{Sport_ID: 1, Sport_name: "Tennis", Sport_Description: "Tennis Sport"})

	return db
}

func TestCreateCourtWithTimeSlots(t *testing.T) {
	DataBase.DB = setupTestDBForCreateCourt()

	courtRequest := map[string]interface{}{
		"Court_Name":     "Court A",
		"Court_Location": "Downtown",
		"Court_Capacity": 10,
		"Court_Status":   1,
		"Sport_name":     "Tennis",
	}

	body, _ := json.Marshal(courtRequest)
	req, err := http.NewRequest("POST", "/CreateCourtWithTimeSlots", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCourtWithTimeSlots)
	handler.ServeHTTP(recorder, req)

	t.Logf("Response Body: %s", recorder.Body.String())

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["message"] != "Court record and TimeSlots added successfully!!" {
		t.Errorf("unexpected response message: %v", response["message"])
	}

	var savedCourt DataBase.Court
	result := DataBase.DB.First(&savedCourt, "court_name = ?", "Court A")

	if result.Error != nil {
		t.Errorf("Court not found in database: %v", result.Error)
	}

	var savedTimeSlots DataBase.Court_TimeSlots
	result = DataBase.DB.First(&savedTimeSlots, "court_id = ?", savedCourt.Court_ID)

	if result.Error != nil {
		t.Errorf("Court timeslots not found in database: %v", result.Error)
	}
}

func TestCreateCourtWithInvalidSport(t *testing.T) {
	var err error
	DataBase.DB, err = setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up test database: %v", err)
	}

	courtRequest := map[string]interface{}{
		"Court_Name":     "Court B",
		"Court_Location": "Downtown",
		"Court_Capacity": 10,
		"Court_Status":   1,
		"Sport_name":     "InvalidSport",
	}

	body, _ := json.Marshal(courtRequest)
	req, _ := http.NewRequest("POST", "/CreateCourtWithTimeSlots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCourtWithTimeSlots)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}

func TestCreateDuplicateCourt(t *testing.T) {
	var err error
	DataBase.DB, err = setupTestDB()
	if err != nil {
		t.Fatalf("failed to set up test database: %v", err)
	}

	courtRequest := map[string]interface{}{
		"Court_Name":     "Court C",
		"Court_Location": "Downtown",
		"Court_Capacity": 10,
		"Court_Status":   1,
		"Sport_name":     "Tennis",
	}

	body, _ := json.Marshal(courtRequest)
	req, _ := http.NewRequest("POST", "/CreateCourtWithTimeSlots", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCourtWithTimeSlots)
	handler.ServeHTTP(recorder, req)

	recorder = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", "/CreateCourtWithTimeSlots", bytes.NewBuffer(body))
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}
