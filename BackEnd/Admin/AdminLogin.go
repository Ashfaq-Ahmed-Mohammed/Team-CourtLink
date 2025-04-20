package Admin

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest

	// Parse JSON body
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Find admin by username and password
	var admin DataBase.Admin
	result := DataBase.DB.Where("Username = ? AND Password = ?", loginReq.Username, loginReq.Password).First(&admin)
	if result.Error != nil || result.RowsAffected == 0 {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Success
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Login successful",
	})
}
