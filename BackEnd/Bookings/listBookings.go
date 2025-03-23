package Bookings

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

func ListBookings(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email query parameter is required", http.StatusBadRequest)
		return
	}

	var customer DataBase.Customer
	err := DataBase.DB.Raw("SELECT * FROM Customer WHERE Email = ?", email).Scan(&customer).Error
	if err != nil || customer.Customer_ID == 0 {
		http.Error(w, "Customer not found", http.StatusNotFound)
		return
	}

	type bookingRaw struct {
		BookingID     uint   `json:"booking_id"`
		CourtName     string `json:"court_name"`
		SportName     string `json:"sport_name"`
		SlotIndex     int    `json:"slot_index"`
		BookingStatus string `json:"booking_status"`
	}

	var bookingsRaw []bookingRaw

	err = DataBase.DB.Raw(`
		SELECT 
			b.Booking_ID as booking_id, 
			c.Court_Name as court_name, 
			s.Sport_name as sport_name, 
			b.Booking_Time as slot_index, 
			b.Booking_Status as booking_status
		FROM Bookings b
		INNER JOIN Court c ON c.Court_ID = b.Court_ID
		INNER JOIN Sport s ON s.Sport_ID = b.Sport_ID
		WHERE b.Customer_ID = ?`,
		customer.Customer_ID).Scan(&bookingsRaw).Error
	if err != nil {
		http.Error(w, "Database error while fetching bookings", http.StatusInternalServerError)
		return
	}

	slots := []string{
		"8-9 AM",
		"9-10 AM",
		"10-11 AM",
		"11-12 AM",
		"12-1 PM",
		"1-2 PM",
		"2-3 PM",
		"3-4 PM",
		"4-5 PM",
		"5-6 PM",
	}

	type BookingResponse struct {
		BookingID     uint   `json:"booking_id"`
		CourtName     string `json:"court_name"`
		SportName     string `json:"sport_name"`
		SlotTime      string `json:"slot_time"`
		BookingStatus string `json:"booking_status"`
	}

	var responseBookings []BookingResponse
	for _, b := range bookingsRaw {
		slotTime := ""
		if b.SlotIndex >= 0 && b.SlotIndex < len(slots) {
			slotTime = slots[b.SlotIndex]
		}
		responseBookings = append(responseBookings, BookingResponse{
			BookingID:     b.BookingID,
			CourtName:     b.CourtName,
			SportName:     b.SportName,
			SlotTime:      slotTime,
			BookingStatus: b.BookingStatus,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBookings)
}
