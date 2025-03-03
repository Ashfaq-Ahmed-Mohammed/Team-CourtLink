package Customer

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var c DataBase.Customer

	// Decode the request body into the Customer struct
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if the customer already exists in the database
	var existingCustomer DataBase.Customer
	result := DataBase.DB.Where("name = ? AND email = ?", c.Name, c.Email).First(&existingCustomer)

	// If customer already exists, return OK response
	if result.RowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Insert the new customer into the database
	if err := DataBase.DB.Create(&c).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send a response confirming the customer was created
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]interface{}{
		"message":  "Customer record added successfully",
		"customer": c,
	}
	json.NewEncoder(w).Encode(response)
}
