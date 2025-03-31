package Court

import (
	"BackEnd/DataBase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDBForDeleteCourt() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&DataBase.Court{}, &DataBase.Court_TimeSlots{})
	return db
}

func TestDeleteCourt(t *testing.T) {

	db := setupTestDBForDeleteCourt()

	DataBase.DB = db

	court := DataBase.Court{
		Court_Name:     "Court A",
		Court_Location: "Downtown",
		Court_Status:   1,
		Sport_id:       1,
	}
	if err := DataBase.DB.Create(&court).Error; err != nil {
		t.Fatalf("failed to create court: %v", err)
	}

	courtTimeSlots := DataBase.Court_TimeSlots{
		Court_ID: court.Court_ID,
	}
	if err := DataBase.DB.Create(&courtTimeSlots).Error; err != nil {
		t.Fatalf("failed to create court time slots: %v", err)
	}

	requestData := map[string]string{
		"Court_Name": "Court A",
	}
	body, _ := json.Marshal(requestData)
	req, err := http.NewRequest("DELETE", "/DeleteCourt", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteCourt)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var response map[string]string
	if err := json.NewDecoder(recorder.Body).Decode(&response); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}
}
