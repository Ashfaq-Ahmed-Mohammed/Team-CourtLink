package Sport

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

func CreateSport(w http.ResponseWriter, r *http.Request) {
	var s DataBase.Sport
	err := json.NewDecoder(r.Body).Decode(&s)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	var existingSport DataBase.Sport
	result := DataBase.DB.Where("Sport_name = ?", s.Sport_name).First(&existingSport)

	if result.RowsAffected > 0 {
		http.Error(w, "The sport record already exists", http.StatusBadRequest)
		return
	}

	if err := DataBase.DB.Create(&s).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Sport record added successfully!!",
		"sport":   s,
	}
	json.NewEncoder(w).Encode(response)
}
