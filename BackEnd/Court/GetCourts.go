package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCourt(w http.ResponseWriter, r *http.Request) {
	var sportselection DataBase.SportSelection
	var sport DataBase.Sport
	var courtStatus uint
	var all_courtstatus []uint
	var courtNames []string
	// var j = 0
	err := json.NewDecoder(r.Body).Decode(&sportselection)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	fmt.Println("Sport Selection:", sportselection.Sport)
	if err := DataBase.DB.Where("Sport_name = ?", sportselection.Sport).First(&sport).Error; err != nil {
		fmt.Println("Sport not found:", err)
		http.Error(w, "Sport not found", http.StatusNotFound)
		return
	}

	if err := DataBase.DB.Model(&DataBase.Court{}).Where("Sport_id = ?", sport.Sport_ID).Pluck("Court_Name", &courtNames).Error; err != nil {
		fmt.Println("Court not found:", err)
		http.Error(w, "Court not found", http.StatusNotFound)
		return
	}

	for i := 0; i < len(courtNames); i++ {
		if err := DataBase.DB.Model(&DataBase.Court{}).Where("Court_Name = ?", courtNames[i]).Pluck("Court_Status", &courtStatus).Error; err != nil {
			fmt.Println("Court Status not found:", err)
			http.Error(w, "Court Status not found", http.StatusNotFound)
			return
		}
		all_courtstatus = append(all_courtstatus, courtStatus)
	}

	var courts []DataBase.CourtAvailability
	for i := 0; i < len(courtNames); i++ {
		court := DataBase.CourtAvailability{
			CourtName:   courtNames[i],
			CourtStatus: all_courtstatus[i],
		}
		courts = append(courts, court)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(courts)
	fmt.Println("Court Info Successful for Sport:", sportselection.Sport)
}
