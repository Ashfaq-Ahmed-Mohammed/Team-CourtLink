package Utils

import (
	"BackEnd/DataBase"
	"log"
)

// ResetTimeSlotsForAvailableCourts resets all slots to 1 for courts that are available (Court_Status = 1)
func ResetTimeSlotsForAvailableCourts() error {
	var courtIDs []uint

	if err := DataBase.DB.
		Model(&DataBase.Court{}).
		Where("court_status = ?", 1).
		Pluck("court_id", &courtIDs).Error; err != nil {
		return err
	}

	if len(courtIDs) == 0 {
		log.Println("No available courts found to reset.")
		return nil
	}

	slotReset := map[string]interface{}{
		"slot_08_09": 1, "slot_09_10": 1, "slot_10_11": 1, "slot_11_12": 1,
		"slot_12_13": 1, "slot_13_14": 1, "slot_14_15": 1, "slot_15_16": 1,
		"slot_16_17": 1, "slot_17_18": 1,
	}

	if err := DataBase.DB.
		Model(&DataBase.Court_TimeSlots{}).
		Where("court_id IN ?", courtIDs).
		Updates(slotReset).Error; err != nil {
		return err
	}

	log.Printf("Reset time slots for %d available court(s).\n", len(courtIDs))
	return nil
}
