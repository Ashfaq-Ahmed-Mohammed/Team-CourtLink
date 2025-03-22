package Court

import (
	"BackEnd/DataBase"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestUpdateCourtSlotandBooking(t *testing.T) {
	// Setup test database.
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to connect to test database: %v", err)
	}
	err = db.AutoMigrate(&DataBase.Customer{}, &DataBase.Sport{}, &DataBase.Court_TimeSlots{}, &DataBase.Bookings{})
	if err != nil {
		t.Fatalf("AutoMigrate failed: %v", err)
	}
	DataBase.DB = db

	// Create a test court timeslot record.
	testTimeSlot := DataBase.Court_TimeSlots{
		Court_ID:   1,
		Court_Name: "Court A",
		Slot_08_09: 1,
		Slot_09_10: 1,
		Slot_10_11: 1,
		Slot_11_12: 1,
		Slot_12_13: 1,
		Slot_13_14: 1,
		Slot_14_15: 1,
		Slot_15_16: 1,
		Slot_16_17: 1,
		Slot_17_18: 1,
	}
	if err := db.Create(&testTimeSlot).Error; err != nil {
		t.Fatalf("Failed to create test timeslot: %v", err)
	}

	// Create a test sport record needed for booking.
	testSport := DataBase.Sport{
		Sport_name: "Tennis",
	}
	if err := db.Create(&testSport).Error; err != nil {
		t.Fatalf("Failed to create test sport: %v", err)
	}

	// Create a test customer record needed for booking.
	testCustomer := DataBase.Customer{
		Customer_ID: 101,
		Name:        "Test Customer",
		Email:       "customer@example.com",
		Contact:     "1234567890",
	}
	if err := db.Create(&testCustomer).Error; err != nil {
		t.Fatalf("Failed to create test customer: %v", err)
	}

	// Construct the JSON request body as a raw string.
	requestBodyStr := `{
		"Court_ID": 1,
		"Slot_Index": 0,
		"Court_Name": "Court A",
		"Sport_name": "Tennis",
		"Customer_email": "customer@example.com"
	}`

	req, err := http.NewRequest("PUT", "/UpdateCourtSlotandBooking", bytes.NewBufferString(requestBodyStr))
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	UpdateCourtSlotandBooking(rr, req)

	resp := rr.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Response Status:", resp.StatusCode)
	fmt.Println("Response Body:", string(body))

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	// Verify the court slot is updated.
	var updatedTimeSlot DataBase.Court_TimeSlots
	if err := db.First(&updatedTimeSlot, "Court_ID = ?", 1).Error; err != nil {
		t.Fatalf("Failed to fetch updated timeslot: %v", err)
	}
	if updatedTimeSlot.Slot_08_09 != 0 {
		t.Errorf("Expected Slot_08_09 to be flipped to 0, got %d", updatedTimeSlot.Slot_08_09)
	}

	// Verify a booking record was created.
	var booking DataBase.Bookings
	if err := db.First(&booking, "Court_ID = ?", 1).Error; err != nil {
		t.Errorf("Expected booking record to be created, got error: %v", err)
	} else {
		if booking.Booking_Status != "booked" {
			t.Errorf("Expected Booking_Status 'booked', got '%s'", booking.Booking_Status)
		}
		// According to the API code, Booking_Time is set to the Slot_Index (0 in this case).
		if booking.Booking_Time != 0 {
			t.Errorf("Expected Booking_Time to be 0, got %v", booking.Booking_Time)
		}
		// Verify that the Customer_ID in the booking matches the test customer.
		if booking.Customer_ID != testCustomer.Customer_ID {
			t.Errorf("Expected Customer_ID to be %d, got %d", testCustomer.Customer_ID, booking.Customer_ID)
		}
	}

	// Allow time for any asynchronous operations (if needed).
	time.Sleep(100 * time.Millisecond)
}
