package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

// DeleteCourt godoc
// @Summary      Delete a court record
// @Description  Deletes a court record from the database based on the court name.
// @Tags         courts
// @Accept       json
// @Produce      json
// @Param        court_name  query     string  true  "Court Name to be deleted"
// @Success      200  {object}  map[string]string  "Court deleted successfully"
// @Failure      400  {object}  map[string]string  "Invalid court name"
// @Failure      404  {object}  map[string]string  "Court not found"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /DeleteCourt [delete]
func DeleteCourt(w http.ResponseWriter, r *http.Request) {
	var requestData struct {
		Court_Name string `json:"Court_Name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var court DataBase.Court
	result := DataBase.DB.Where("Court_Name = ?", requestData.Court_Name).First(&court)
	if result.RowsAffected == 0 {
		http.Error(w, "Court not found", http.StatusNotFound)
		return
	}

	if err := DataBase.DB.Where("Court_ID = ?", court.Court_ID).Delete(&DataBase.Court_TimeSlots{}).Error; err != nil {
		http.Error(w, "Failed to delete court time slots", http.StatusInternalServerError)
		return
	}

	if err := DataBase.DB.Delete(&court).Error; err != nil {
		http.Error(w, "Failed to delete court", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Court deleted successfully"}
	json.NewEncoder(w).Encode(response)
}
