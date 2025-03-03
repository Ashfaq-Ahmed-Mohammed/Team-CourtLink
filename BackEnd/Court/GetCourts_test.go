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

	// Seed test data
	testSport := DataBase.Sport{Sport_ID: 1, Sport_name: "tennis"}
	db.Create(&testSport)

	testCourt := DataBase.Court{Court_ID: 1, Court_Name: "Court A", Court_Status: 1, Sport_id: 1}
	db.Create(&testCourt)

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

	tests := []struct {
		name           string
		query          string
		expectedCode   int
		expectedCourts int
	}{
		{"Valid sport with court", "/getCourts?sport=tennis", http.StatusOK, 1},
		{"Invalid sport", "/getCourts?sport=badminton", http.StatusNotFound, 0},
		{"Missing sport param", "/getCourts", http.StatusBadRequest, 0},
		{"Sport with no courts", "/getCourts?sport=squash", http.StatusNotFound, 0},
		{"Courts exist but no timeslots", "/getCourts?sport=football", http.StatusNotFound, 0},
	}

	// Create a sport with no courts
	testSportNoCourts := DataBase.Sport{Sport_ID: 2, Sport_name: "squash"}
	db.Create(&testSportNoCourts)

	// Create a sport with a court but no timeslots
	testSportNoSlots := DataBase.Sport{Sport_ID: 3, Sport_name: "football"}
	db.Create(&testSportNoSlots)
	testCourtNoSlots := DataBase.Court{Court_ID: 2, Court_Name: "Court B", Court_Status: 1, Sport_id: 3}
	db.Create(&testCourtNoSlots)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.query, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			w := httptest.NewRecorder()
			GetCourt(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tc.expectedCode {
				t.Errorf("Expected status %d, got %d", tc.expectedCode, resp.StatusCode)
			}

			if resp.StatusCode == http.StatusOK {
				var courts []DataBase.CourtAvailability
				if err := json.NewDecoder(resp.Body).Decode(&courts); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}

				if len(courts) != tc.expectedCourts {
					t.Errorf("Expected %d courts, got %d", tc.expectedCourts, len(courts))
				}

				if tc.expectedCourts > 0 && (courts[0].CourtID != 1 || !strings.Contains(courts[0].CourtName, "Court A")) {
					t.Errorf("Unexpected court data: %v", courts[0])
				}
			}
		})
	}
}
