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

func TestCancelBookingandUpdateSlot(t *testing.T) {
	// Initialize a fresh test DB and assign it to the global variable.
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	DataBase.DB = db

	// Insert test data.
	// Create a Court_TimeSlots record for Court_ID = 101.
	timeSlots := DataBase.Court_TimeSlots{
		Court_ID:   102,
		Court_Name: "Court B", // Provide a unique court name to avoid UNIQUE constraint error.
		Slot_08_09: 0,
		Slot_09_10: 0,
		Slot_10_11: 0,
		Slot_11_12: 0, // This slot corresponds to index 3.
		Slot_12_13: 0,
		Slot_13_14: 0,
		Slot_14_15: 0,
		Slot_15_16: 0,
		Slot_16_17: 0,
		Slot_17_18: 0,
	}
	if err := DataBase.DB.Create(&timeSlots).Error; err != nil {
		t.Fatalf("failed to create Court_TimeSlots record: %v", err)
	}

	// Create a booking record with Booking_ID=1, Court_ID=101, and Booking_Time=3 (maps to "slot_11_12").
	booking := DataBase.Bookings{
		Booking_ID:     2,
		Customer_ID:    122,
		Court_ID:       102,
		Sport_ID:       122,
		Booking_Status: "Confirmed",
		Booking_Time:   3, // corresponds to "slot_11_12"
	}
	if err := DataBase.DB.Create(&booking).Error; err != nil {
		t.Fatalf("failed to create Booking record: %v", err)
	}

	// Prepare the cancel request payload as a raw JSON string.
	reqBody := `{"Booking_ID":2}`

	// Create the HTTP request.
	req, err := http.NewRequest("PUT", "/CancelBookingandUpdateSlot", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Execute the handler.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CancelBookingandUpdateSlot)
	handler.ServeHTTP(rr, req)

	// Check the response status code.
	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	// Verify the response message.
	expectedMessage := fmt.Sprintf("Booking cancelled and slot updated successfully for Booking_ID: %d", 2)
	if rr.Body.String() != expectedMessage {
		t.Errorf("expected response message %q, got %q", expectedMessage, rr.Body.String())
	}

	// Confirm that the booking status was updated to "Cancelled" in the database.
	var updatedBooking DataBase.Bookings
	if err := DataBase.DB.First(&updatedBooking, 2).Error; err != nil {
		t.Fatalf("failed to query updated booking: %v", err)
	}
	if updatedBooking.Booking_Status != "Cancelled" {
		t.Errorf("expected booking status to be 'Cancelled', got %q", updatedBooking.Booking_Status)
	}

	// Verify that the correct timeslot (slot_11_12) was updated to 1.
	var updatedTimeSlots DataBase.Court_TimeSlots
	if err := DataBase.DB.Where("Court_ID = ?", 102).First(&updatedTimeSlots).Error; err != nil {
		t.Fatalf("failed to query updated time slots: %v", err)
	}
	if updatedTimeSlots.Slot_11_12 != 1 {
		t.Errorf("expected Slot_11_12 to be 1, got %d", updatedTimeSlots.Slot_11_12)
	}
}

func TestResetCourtSlotsHandler(t *testing.T) {

	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("failed to setup test database: %v", err)
	}
	DataBase.DB = db

	court := DataBase.Court{
		Court_Name:     "Court C",
		Court_Location: "Test Location",
		Court_Status:   1,
		Sport_id:       1,
	}
	if err := DataBase.DB.Create(&court).Error; err != nil {
		t.Fatalf("failed to create Court: %v", err)
	}

	timeSlots := DataBase.Court_TimeSlots{
		Court_ID:   court.Court_ID,
		Court_Name: court.Court_Name,
		Slot_08_09: 0,
		Slot_09_10: 0,
		Slot_10_11: 0,
		Slot_11_12: 0,
		Slot_12_13: 0,
		Slot_13_14: 0,
		Slot_14_15: 0,
		Slot_15_16: 0,
		Slot_16_17: 0,
		Slot_17_18: 0,
	}
	if err := DataBase.DB.Create(&timeSlots).Error; err != nil {
		t.Fatalf("failed to create Court_TimeSlots: %v", err)
	}

	req, err := http.NewRequest("PUT", "/resetCourtSlots", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ResetCourtSlotsHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	expectedMessage := "Court slots reset successfully!"
	if rr.Body.String() != expectedMessage {
		t.Errorf("expected response message %q, got %q", expectedMessage, rr.Body.String())
	}

	var updatedTimeSlots DataBase.Court_TimeSlots
	if err := DataBase.DB.First(&updatedTimeSlots, "court_id = ?", court.Court_ID).Error; err != nil {
		t.Fatalf("failed to fetch updated Court_TimeSlots: %v", err)
	}

	slots := []int{
		updatedTimeSlots.Slot_08_09, updatedTimeSlots.Slot_09_10, updatedTimeSlots.Slot_10_11,
		updatedTimeSlots.Slot_11_12, updatedTimeSlots.Slot_12_13, updatedTimeSlots.Slot_13_14,
		updatedTimeSlots.Slot_14_15, updatedTimeSlots.Slot_15_16, updatedTimeSlots.Slot_16_17,
		updatedTimeSlots.Slot_17_18,
	}

	for idx, slot := range slots {
		if slot != 1 {
			t.Errorf("expected slot %d to be 1 (available), got %d", idx, slot)
		}
	}
}
