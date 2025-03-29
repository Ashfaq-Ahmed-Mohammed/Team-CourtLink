package Bookings

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

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// Migrate the schema.
	db.AutoMigrate(&DataBase.Customer{}, &DataBase.Sport{}, &DataBase.Court{}, &DataBase.Bookings{})

	// Insert test data.
	db.Create(&DataBase.Customer{
		Customer_ID: 122,
		Name:        "John Doe",
		Email:       "john@example.com",
		Contact:     "1234567890",
	})
	db.Create(&DataBase.Sport{
		Sport_ID:          122,
		Sport_name:        "Tennis",
		Sport_Description: "Tennis Sport",
	})
	db.Create(&DataBase.Court{
		Court_ID:       122,
		Court_Name:     "Court A",
		Court_Location: "Downtown",
		Court_Status:   1,
		Sport_id:       122,
	})
	// Insert a booking record with the correct Booking_Time.
	db.Create(&DataBase.Bookings{
		Booking_ID:     1,
		Customer_ID:    122,
		Court_ID:       122,
		Sport_ID:       122,
		Booking_Status: "Confirmed",
		Booking_Time:   2, // This corresponds to "10-11 AM"
	})

	return db
}

func TestCreateBooking(t *testing.T) {
	DataBase.DB = setupTestDB()

	bookingRequest := map[string]interface{}{
		"Customer_ID":    122,
		"Sport_ID":       122,
		"Court_ID":       122,
		"Booking_Status": "Confirmed",
	}

	body, _ := json.Marshal(bookingRequest)
	req, err := http.NewRequest("POST", "/CreateBooking", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateBooking)
	handler.ServeHTTP(recorder, req)

	t.Logf("Response Body: %s", recorder.Body.String())

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)

	if response["message"] != "Booking record added successfully" {
		t.Errorf("unexpected response message: %v", response["message"])
	}

	var savedBooking DataBase.Bookings
	result := DataBase.DB.First(&savedBooking, "customer_id = ? AND court_id = ?", 122, 122)

	t.Logf("Database Query Result: %+v", savedBooking)

	if result.Error != nil {
		t.Errorf("Booking not found in database: %v", result.Error)
	} else if savedBooking.Customer_ID != 122 || savedBooking.Court_ID != 122 {
		t.Errorf("Booking data mismatch: got %+v", savedBooking)
	}
}
