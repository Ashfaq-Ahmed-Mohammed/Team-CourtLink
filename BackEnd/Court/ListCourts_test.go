package Court

import (
	"BackEnd/DataBase"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupListCourtsTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&DataBase.Sport{}, &DataBase.Court{})

	db.Create(&DataBase.Sport{
		Sport_name:        "Football",
		Sport_Description: "A popular team sport",
	})
	db.Create(&DataBase.Court{
		Court_Name:     "Court A",
		Court_Location: "Downtown",
		Court_Status:   1,
		Sport_id:       1,
	})

	return db
}

func TestListCourts(t *testing.T) {
	DataBase.DB = setupListCourtsTestDB()

	req, err := http.NewRequest("GET", "/ListCourts", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListCourts)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var courts []CourtData
	err = json.Unmarshal(recorder.Body.Bytes(), &courts)
	if err != nil {
		t.Errorf("Failed to unmarshal response body: %v", err)
	}

	if len(courts) == 0 {
		t.Errorf("Expected courts to be listed, but got an empty response")
	}

	if courts[0].CourtName != "Court A" || courts[0].SportName != "Football" {
		t.Errorf("unexpected court data: %+v", courts)
	}
}
