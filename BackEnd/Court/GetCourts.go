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
	var court_TimeSlot DataBase.Court_TimeSlots
	var courts []DataBase.CourtAvailability
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

	for i := 0; i < len(courtNames); i++ {
		if err := DataBase.DB.Model(&DataBase.Court_TimeSlots{}).Where("Court_Name = ?", courtNames[i]).Find(&court_TimeSlot).Error; err != nil {
			fmt.Println("Court TimeSlots not found:", err)
			http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
			return
		}
		court := DataBase.CourtAvailability{
			CourtName:   courtNames[i],
			CourtStatus: all_courtstatus[i],
			Slots:       []int{court_TimeSlot.Slot_8_9, court_TimeSlot.Slot_9_10, court_TimeSlot.Slot_10_11, court_TimeSlot.Slot_11_12, court_TimeSlot.Slot_12_13, court_TimeSlot.Slot_13_14, court_TimeSlot.Slot_14_15, court_TimeSlot.Slot_15_16, court_TimeSlot.Slot_16_17, court_TimeSlot.Slot_17_18},
		}
		courts = append(courts, court)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(courts)
	fmt.Println("Court Info Successful for Sport:", sportselection.Sport)
}
