package Sport

import (
	"BackEnd/DataBase"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListSports(w http.ResponseWriter, r *http.Request) {
	var sports []string

	if err := DataBase.DB.Model(&DataBase.Sport{}).Pluck("Sport_name", &sports).Error; err != nil {
		fmt.Println("Failed to fetch sports:", err)
		http.Error(w, "Failed to fetch sports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sports)

	fmt.Println("ListSports API called successfully")
}
