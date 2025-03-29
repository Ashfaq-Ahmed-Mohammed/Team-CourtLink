package Bookings

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListBookings(t *testing.T) {
	DataBase.DB = setupTestDB()

	req, err := http.NewRequest("GET", "/ListBookings?email=john@example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	handler := http.HandlerFunc(ListBookings)
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}

	var responseBookings []struct {
		BookingID     uint   `json:"booking_id"`
		CourtName     string `json:"court_name"`
		SportName     string `json:"sport_name"`
		SlotTime      string `json:"slot_time"`
		BookingStatus string `json:"booking_status"`
	}
	if err := json.Unmarshal(recorder.Body.Bytes(), &responseBookings); err != nil {
		t.Errorf("error unmarshalling response: %v", err)
	}

	if len(responseBookings) != 1 {
		t.Errorf("expected 1 booking, got %d", len(responseBookings))
	}

	booking := responseBookings[0]
	if booking.BookingID != 1 {
		t.Errorf("expected BookingID %d, got %d", 1, booking.BookingID)
	}
	if booking.CourtName != "Court A" {
		t.Errorf("expected CourtName %s, got %s", "Court A", booking.CourtName)
	}
	if booking.SportName != "Tennis" {
		t.Errorf("expected SportName %s, got %s", "Tennis", booking.SportName)
	}
	if booking.SlotTime != "10-11 AM" {
		t.Errorf("expected SlotTime %s, got %s", "10-11 AM", booking.SlotTime)
	}
	if booking.BookingStatus != "Confirmed" {
		t.Errorf("expected BookingStatus %s, got %s", "Confirmed", booking.BookingStatus)
	}
}
