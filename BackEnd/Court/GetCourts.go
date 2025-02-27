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
	type CourtInfo struct {
		CourtID     uint   `gorm:"column:Court_ID"`
		CourtName   string `gorm:"column:Court_Name"`
		CourtStatus uint   `gorm:"column:Court_Status"`
	}
	var courtData []CourtInfo
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

	if err := DataBase.DB.Model(&DataBase.Court{}).
		Select("Court_ID, Court_Name, Court_Status").
		Where("Sport_id = ?", sport.Sport_ID).
		Find(&courtData).Error; err != nil {
		fmt.Println("Court not found:", err)
		http.Error(w, "Court not found", http.StatusNotFound)
		return
	}

	// Extract court IDs for time slot retrieval
	var courtIDs []uint
	for _, court := range courtData {
		courtIDs = append(courtIDs, court.CourtID)
	}

	// Retrieve all time slots using **Court_ID** in a **single batch query**
	var courtTimeSlots []DataBase.Court_TimeSlots
	if err := DataBase.DB.
		Where("Court_ID IN (?)", courtIDs).
		Find(&courtTimeSlots).Error; err != nil {
		fmt.Println("Court TimeSlots not found:", err)
		http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
		return
	}

	// Create a map of Court_ID -> Court_TimeSlots for quick lookup
	courtTimeSlotMap := make(map[uint]DataBase.Court_TimeSlots)
	for _, slot := range courtTimeSlots {
		courtTimeSlotMap[slot.Court_ID] = slot
	}

	// Process data efficiently and return **Court_ID, Court_Name, Court_Status, and Slots**
	var courts []DataBase.CourtAvailability
	for _, court := range courtData {
		timeSlot, exists := courtTimeSlotMap[court.CourtID]
		if !exists {
			fmt.Println("No time slots found for Court ID:", court.CourtID)
			http.Error(w, "Court TimeSlots not found", http.StatusNotFound)
			return
		}

		courtAvailability := DataBase.CourtAvailability{
			CourtID:     court.CourtID,
			CourtName:   court.CourtName,
			CourtStatus: uint(court.CourtStatus),
			Slots:       []int{timeSlot.Slot_08_09, timeSlot.Slot_09_10, timeSlot.Slot_10_11, timeSlot.Slot_11_12, timeSlot.Slot_12_13, timeSlot.Slot_13_14, timeSlot.Slot_14_15, timeSlot.Slot_15_16, timeSlot.Slot_16_17, timeSlot.Slot_17_18},
		}
		courts = append(courts, courtAvailability)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(courts)
	fmt.Println("Court Info Successful for Sport:", sportselection.Sport)
}
