package Bookings

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

// CreateBooking creates a new booking after validating customer, sport, and court.
// @Summary Create a new booking
// @Description Creates a new booking after validating the existence of the customer, sport, and court.
// @Tags bookings
// @Accept json
// @Produce json
// @Param  booking body DataBase.Bookings true "Booking data"
// @Success 201 {object} map[string]interface{} "Booking record added successfully"
// @Failure 400 {object} DataBase.ErrorResponse "Invalid request body"
// @Failure 404 {object} DataBase.ErrorResponse "Customer, sport, or court not found"
// @Failure 500 {object} DataBase.ErrorResponse "Internal server error"
// @Router /CreateBooking [post]
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	var b DataBase.Bookings

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var customer DataBase.Customer
	var sport DataBase.Sport
	var court DataBase.Court

	if DataBase.DB.First(&customer, b.Customer_ID).RowsAffected == 0 {
		http.Error(w, "Customer is not found", http.StatusNotFound)
		return
	}
	if DataBase.DB.First(&sport, b.Sport_ID).RowsAffected == 0 {
		http.Error(w, "Sport is not found", http.StatusNotFound)
		return
	}
	if DataBase.DB.First(&court, b.Court_ID).RowsAffected == 0 {
		http.Error(w, "Court is not found", http.StatusNotFound)
		return
	}

	if err := DataBase.DB.Create(&b).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message": "Booking record added successfully",
		"booking": b,
	}
	json.NewEncoder(w).Encode(response)
}
