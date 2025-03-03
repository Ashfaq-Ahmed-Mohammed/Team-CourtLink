package Customer

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

// CreateCustomer handles customer creation.
// @Summary Create a new customer
// @Description Adds a new customer to the database if they do not already exist.
// @Tags customers
// @Accept json
// @Produce json
// @Param customer body DataBase.Customer true "Customer data"
// @Success 201 {object} map[string]interface{} "Customer record added successfully"
// @Success 200 "Customer already exists"
// @Failure 400 "Invalid request body"
// @Failure 500 "Internal server error"
// @Router /customer [post]
func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	var c DataBase.Customer

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	var existingCustomer DataBase.Customer
	result := DataBase.DB.Where("name = ? AND email = ?", c.Name, c.Email).First(&existingCustomer)

	if result.RowsAffected > 0 {
		w.WriteHeader(http.StatusOK)
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
