package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&DataBase.Sport{}, &DataBase.Court{}, &DataBase.Court_TimeSlots{})

	return db, nil
}

func TestGetCourt(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	DataBase.DB = db

	// Create a test sport
	testSport := DataBase.Sport{Sport_ID: 1, Sport_name: "tennis"}
	db.Create(&testSport)

	// Create a test court
	testCourt := DataBase.Court{Court_ID: 1, Court_Name: "Court A", Court_Status: 1, Sport_id: 1}
	db.Create(&testCourt)

	// Create a test time slot
	testTimeSlot := DataBase.Court_TimeSlots{
		Court_ID:   1,
		Slot_08_09: 1,
		Slot_09_10: 0,
		Slot_10_11: 1,
		Slot_11_12: 0,
		Slot_12_13: 1,
		Slot_13_14: 0,
		Slot_14_15: 1,
		Slot_15_16: 0,
		Slot_16_17: 1,
		Slot_17_18: 0,
	}
	db.Create(&testTimeSlot)

	r, err := http.NewRequest("GET", "/getCourts?sport=tennis", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	w := httptest.NewRecorder()
	GetCourt(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Error: %v", resp.Status)
	}

	var courts []DataBase.CourtAvailability
	if err := json.NewDecoder(resp.Body).Decode(&courts); err != nil {
		t.Errorf("Error: %v", err)
	}

	if len(courts) == 0 {
		t.Errorf("No Courts displayed")
	}

	if courts[0].CourtID != 1 || !strings.Contains(courts[0].CourtName, "Court A") {
		t.Errorf("Unexpected court data: %v", courts[0])
	}
}
