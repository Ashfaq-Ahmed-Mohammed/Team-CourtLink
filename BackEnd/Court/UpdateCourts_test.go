package Court

import (
	"BackEnd/DataBase"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateCourtSlot(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}

	DataBase.DB = db

	// Create a test court time slot
	testTimeSlot := DataBase.Court_TimeSlots{
		Court_ID:   1,
		Court_Name: "Court A",
		Slot_08_09: 1,
		Slot_09_10: 1,
		Slot_10_11: 0,
	}
	db.Create(&testTimeSlot)

	// Create an update request
	updateRequest := DataBase.CourtUpdate{
		Court_ID:    1,
		Slot_Index:  0,
		Court_Name:  "Court A",
		Sport_name:  "Tennis",
		Customer_ID: 101,
	}
	requestBody, err := json.Marshal(updateRequest)

	if err != nil {
		t.Fatalf("Could not marshal request: %v", err)
	}
	r, err := http.NewRequest("PUT", "/UpdateCourtSlot", bytes.NewBuffer(requestBody))
	r.Header.Set("Content-Type", "application/json")

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	w := httptest.NewRecorder()
	UpdateCourtSlot(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Status:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	// Check if slot value flipped
	var updatedTimeSlot DataBase.Court_TimeSlots
	db.First(&updatedTimeSlot, "Court_ID = ?", 1)
	if updatedTimeSlot.Slot_08_09 != 0 {
		t.Errorf("Expected Slot_08_09 to be flipped, but got %d", updatedTimeSlot.Slot_08_09)
	}
}
