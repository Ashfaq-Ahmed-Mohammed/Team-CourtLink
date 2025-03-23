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

	type BookingResponse struct {
		BookingID     uint   `json:"booking_id"`
		CourtName     string `json:"court_name"`
		SportName     string `json:"sport_name"`
		SlotIndex     int    `json:"slot_index"`
		BookingStatus string `json:"booking_status"`
	}

	var booked []BookingResponse
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
		WHERE b.Customer_ID = ? AND b.Booking_Status = ?`,
		customer.Customer_ID, "booked").Scan(&booked).Error
	if err != nil {
		http.Error(w, "Database error while fetching booked bookings", http.StatusInternalServerError)
		return
	}

	var cancelled []BookingResponse
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
		WHERE b.Customer_ID = ? AND b.Booking_Status = ?`,
		customer.Customer_ID, "cancelled").Scan(&cancelled).Error
	if err != nil {
		http.Error(w, "Database error while fetching cancelled bookings", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"booked":    booked,
		"cancelled": cancelled,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
