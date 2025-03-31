package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"fmt"
	"net/http"
)

type CourtData struct {
	CourtName string `json:"court_name"`
	SportName string `json:"sport_name"`
}

// ListCourts godoc
// @Summary      List all courts with their associated sports
// @Description  Retrieves a list of all courts along with the corresponding sport names.
// @Tags         courts
// @Accept       json
// @Produce      json
// @Success      200    {array}   CourtData  "List of courts and their associated sports"  example([{"court_name": "Court A", "sport_name": "Tennis"}, {"court_name": "Court B", "sport_name": "Basketball"}])
// @Failure      500    {string}  string  "Database error while fetching courts"
// @Router       /ListCourts [get]
func ListCourts(w http.ResponseWriter, r *http.Request) {
	var courts []CourtData

	err := DataBase.DB.Table("Court").
		Select("Court.Court_Name as CourtName, Sport.Sport_name as SportName").
		Joins("JOIN Sport ON Sport.Sport_ID = Court.Sport_id").
		Scan(&courts).Error

	if err != nil {
		fmt.Println("Failed to fetch courts:", err)
		http.Error(w, "Failed to fetch courts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courts)

	fmt.Println("ListCourts API called successfully")
}
