package Utils

import (
	"BackEnd/DataBase"
	"log"
)

// ResetTimeSlotsForAvailableCourts resets all slots to 1 for courts that are available (Court_Status = 1)
func ResetTimeSlotsForAvailableCourts(courtName string) error {
	const (
		AvailableStatus = 1
		SlotAvailable   = 1
	)

	var courtIDs []uint
	db := DataBase.DB.Model(&DataBase.Court{}).Where("court_status = ?", AvailableStatus)

	// If a court name is provided, filter by that name
	if courtName != "" {
		db = db.Where("court_name = ?", courtName)
	}

	// Get the court IDs that match the condition(s)
	if err := db.Pluck("court_id", &courtIDs).Error; err != nil {
		return err
	}

	if len(courtIDs) == 0 {
		if courtName != "" {
			log.Printf("No available court found with name: %s\n", courtName)
		} else {
			log.Println("No available courts found to reset.")
		}
		return nil
	}

	slotReset := map[string]interface{}{
		"slot_08_09": SlotAvailable, "slot_09_10": SlotAvailable,
		"slot_10_11": SlotAvailable, "slot_11_12": SlotAvailable,
		"slot_12_13": SlotAvailable, "slot_13_14": SlotAvailable,
		"slot_14_15": SlotAvailable, "slot_15_16": SlotAvailable,
		"slot_16_17": SlotAvailable, "slot_17_18": SlotAvailable,
	}

	if err := DataBase.DB.
		Model(&DataBase.Court_TimeSlots{}).
		Where("court_id IN ?", courtIDs).
		Updates(slotReset).Error; err != nil {
		return err
	}

	if courtName != "" {
		log.Printf("Reset time slots for court: %s\n", courtName)
	} else {
		log.Printf("Reset time slots for %d available court(s).\n", len(courtIDs))
	}

	return nil
}
