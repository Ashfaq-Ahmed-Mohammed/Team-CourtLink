package Customer

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var c DataBase.Customer

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {

		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := DataBase.DB.Create(&c).Error; err != nil {

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message":  "Customer record added successfully",
		"customer": c,
	}
	json.NewEncoder(w).Encode(response)
}
