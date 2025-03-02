package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

// UpdateCourtSlot updates the availability status of a specific court slot.
//
// @Summary Update court slot status
// @Description Toggles the availability of a court time slot. If the slot is booked, it is freed; if it is free, it is booked.
// @Tags courts
// @Accept  json
// @Produce  json
// @Param  updateRequest body DataBase.CourtUpdate true "Court slot update request"
// @Success 200 {string} string "Slot updated successfully for Court_ID: {Court_ID}, Slot_Index: {Slot_Index}"
// @Failure 400 {object} DataBase.ErrorResponse "Invalid request body or Slot_Index out of range"
// @Failure 404 {object} DataBase.ErrorResponse "Court time slots not found"
// @Failure 500 {object} DataBase.ErrorResponse "Database error or failed to update slot"
// @Router /UpdateCourtSlot [put]
func UpdateCourtSlot(w http.ResponseWriter, r *http.Request) {
	var updateRequest DataBase.CourtUpdate

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var timeSlot DataBase.Court_TimeSlots
	if err := DataBase.DB.Model(&DataBase.Court_TimeSlots{}).Where("Court_ID = ?", updateRequest.Court_ID).First(&timeSlot).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	slotFields := []string{
		"slot_08_09", "slot_09_10", "slot_10_11", "slot_11_12",
		"slot_12_13", "slot_13_14", "slot_14_15", "slot_15_16",
		"slot_16_17", "slot_17_18",
	}

	if updateRequest.Slot_Index < 0 || updateRequest.Slot_Index >= len(slotFields) {
		http.Error(w, "Invalid Slot_Index", http.StatusBadRequest)
		return
	}

	fieldName := slotFields[updateRequest.Slot_Index]
	updateQuery := fmt.Sprintf("%s = CASE %s WHEN 0 THEN 1 ELSE 0 END", fieldName, fieldName)

	if err := DataBase.DB.Model(&DataBase.Court_TimeSlots{}).
		Where("Court_ID = ?", updateRequest.Court_ID).
		UpdateColumn(fieldName, gorm.Expr(updateQuery)).Error; err != nil {
		http.Error(w, "Failed to update slot", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Slot updated successfully for Court_ID: %d, Slot_Index: %d", updateRequest.Court_ID, updateRequest.Slot_Index)
}
