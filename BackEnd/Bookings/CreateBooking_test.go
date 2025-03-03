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

func setupTestDB() {

	DataBase.DB, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	DataBase.DB.AutoMigrate(&DataBase.Customer{}, &DataBase.Sport{}, &DataBase.Court{}, &DataBase.Bookings{})

	DataBase.DB.Create(&DataBase.Customer{Customer_ID: 122, Name: "John Doe", Email: "john@example.com", Contact: "1234567890"})
	DataBase.DB.Create(&DataBase.Sport{Sport_ID: 122, Sport_name: "Tennis", Sport_Description: "Tennis Sport"})
	DataBase.DB.Create(&DataBase.Court{Court_ID: 122, Court_Name: "Court A", Court_Location: "Downtown", Court_Status: 1, Sport_id: 1})
}

func TestCreateBooking(t *testing.T) {
	setupTestDB()

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
