package Court

import (
	"BackEnd/DataBase"
	"BackEnd/Utils"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// UpdateCourtSlotandBooking updates the availability status of a specific court slot and creates a corresponding booking record.
//
// @Summary Update court slot and create booking
// @Description Toggles the availability of a court time slot and, based on the provided customer email and sport name,
//
//	creates a booking record if the slot update is successful. Both operations are executed within a single transaction.
//
// @Tags courts
// @Accept json
// @Produce json
// @Param updateRequest body DataBase.CourtUpdate true "Court slot update request including Customer_email and Sport_name"
// @Success 200 {string} string "Slot updated and booking created successfully for Court_ID: {Court_ID}, Slot_Index: {Slot_Index}"
// @Failure 400 {object} DataBase.ErrorResponse "Invalid request body or Slot_Index out of range"
// @Failure 404 {object} DataBase.ErrorResponse "Court time slots, Customer, or Sport not found"
// @Failure 500 {object} DataBase.ErrorResponse "Database error or failed to update slot/booking"
// @Router /UpdateCourtSlotandBooking [put]
func UpdateCourtSlotandBooking(w http.ResponseWriter, r *http.Request) {
	var updateRequest DataBase.CourtUpdate

	// Decode JSON request.
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Start a transaction.
	tx := DataBase.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	// ----- Step 1: Update the Court Slot -----
	var timeSlot DataBase.Court_TimeSlots
	if err := tx.Model(&DataBase.Court_TimeSlots{}).
		Where("Court_ID = ?", updateRequest.Court_ID).
		First(&timeSlot).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	slotFields := []string{
		"slot_08_09", "slot_09_10", "slot_10_11", "slot_11_12",
		"slot_12_13", "slot_13_14", "slot_14_15", "slot_15_16",
		"slot_16_17", "slot_17_18",
	}

	if updateRequest.Slot_Index < 0 || updateRequest.Slot_Index >= len(slotFields) {
		tx.Rollback()
		http.Error(w, "Invalid Slot_Index", http.StatusBadRequest)
		return
	}

	fieldName := slotFields[updateRequest.Slot_Index]

	if err := tx.Model(&DataBase.Court_TimeSlots{}).
		Where("Court_ID = ?", updateRequest.Court_ID).
		UpdateColumn(fieldName, 0).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update slot", http.StatusInternalServerError)
		return
	}

	var customer DataBase.Customer
	if err := tx.Where("email = ?", updateRequest.Customer_email).First(&customer).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Customer not found in the database", http.StatusNotFound)
		} else {
			http.Error(w, "Failed to fetch customer", http.StatusInternalServerError)
		}
		return
	}

	// Lookup the sport by name.
	var sport DataBase.Sport
	if tx.Where("Sport_name = ?", updateRequest.Sport_name).First(&sport).RowsAffected == 0 {
		tx.Rollback()
		http.Error(w, "Sport not found", http.StatusNotFound)
		return
	}

	booking := DataBase.Bookings{
		Customer_ID:    customer.Customer_ID,
		Sport_ID:       sport.Sport_ID,
		Court_ID:       updateRequest.Court_ID,
		Booking_Status: "booked",
		Booking_Time:   updateRequest.Slot_Index,
	}

	if err := tx.Create(&booking).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to create booking", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Slot updated and booking created successfully for Court_ID: %d, Slot_Index: %d", updateRequest.Court_ID, updateRequest.Slot_Index)

}

// CancelBookingandUpdateSlot godoc
// @Summary      Cancel a booking and update court time slot
// @Description  Cancels a booking by updating its status to "Cancelled" and marks the corresponding time slot (based on Booking_Time) as available (sets it to 1) in the Court_TimeSlots record.
// @Tags         courts
// @Accept       json
// @Produce      plain
// @Param        cancelRequest  body      DataBase.CancelRequest  true  "Cancel Booking Request"  example({"Booking_ID": 123})
// @Success      200            {string}  string  "Booking cancelled and slot updated successfully for Booking_ID: 123"
// @Failure      400            {string}  string  "Invalid request body or Invalid Slot_Index"
// @Failure      404            {string}  string  "Booking not found or Court TimeSlots not found"
// @Failure      500            {string}  string  "Failed to start transaction, database error, or transaction commit failed"
// @Router       /CancelBookingandUpdateSlot [put]
func CancelBookingandUpdateSlot(w http.ResponseWriter, r *http.Request) {
	var cancelRequest DataBase.CancelRequest

	if err := json.NewDecoder(r.Body).Decode(&cancelRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	tx := DataBase.DB.Begin()
	if tx.Error != nil {
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}

	var booking DataBase.Bookings
	if err := tx.First(&booking, cancelRequest.Booking_ID).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Booking not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error while fetching booking", http.StatusInternalServerError)
		}
		return
	}

	if err := tx.Model(&booking).UpdateColumn("Booking_Status", "Cancelled").Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update the booking status", http.StatusInternalServerError)
		return
	}

	var timeSlot DataBase.Court_TimeSlots
	if err := tx.Model(&DataBase.Court_TimeSlots{}).
		Where("Court_ID = ?", booking.Court_ID).
		First(&timeSlot).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error while fetching timeslot", http.StatusInternalServerError)
		}
		return
	}

	slotFields := []string{
		"slot_08_09", "slot_09_10", "slot_10_11", "slot_11_12",
		"slot_12_13", "slot_13_14", "slot_14_15", "slot_15_16",
		"slot_16_17", "slot_17_18",
	}
	if booking.Booking_Time < 0 || booking.Booking_Time >= len(slotFields) {
		tx.Rollback()
		http.Error(w, "Invalid Slot_Index", http.StatusBadRequest)
		return
	}
	fieldName := slotFields[booking.Booking_Time]

	if err := tx.Model(&DataBase.Court_TimeSlots{}).
		Where("Court_ID = ?", booking.Court_ID).
		UpdateColumn(fieldName, 1).Error; err != nil {
		tx.Rollback()
		http.Error(w, "Failed to update slot", http.StatusInternalServerError)
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		http.Error(w, "Transaction commit failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Booking cancelled and slot updated successfully for Booking_ID: %d", cancelRequest.Booking_ID)
}

func ResetCourtSlotsHandler(w http.ResponseWriter, r *http.Request) {
	err := Utils.ResetTimeSlotsForAvailableCourts()
	if err != nil {
		http.Error(w, "Failed to reset court slots", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Court slots reset successfully!"))
}
