package Customer

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

func setupTestDB() {
	DataBase.DB, _ = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	DataBase.DB.AutoMigrate(&DataBase.Customer{})
}

func TestCreateCustomer(t *testing.T) {
	setupTestDB()

	customer := DataBase.Customer{
		Name:    "Rohi B",
		Email:   "rohb@example.com",
		Contact: "1234567890",
	}

	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/Customer", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	if response["message"] != "Customer record added successfully" {
		t.Errorf("response message unexpected: %v", response["message"])
	}
}

func TestCreateDuplicateCustomer(t *testing.T) {
	setupTestDB()

	customer := DataBase.Customer{
		Name:    "Jane Doe",
		Email:   "janedoe@example.com",
		Contact: "0987654321",
	}
	DataBase.DB.Create(&customer)

	body, _ := json.Marshal(customer)
	req, _ := http.NewRequest("POST", "/Customer", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}
}

func TestCreateCustomerInvalidRequest(t *testing.T) {
	setupTestDB()

	req, _ := http.NewRequest("POST", "/Customer", bytes.NewBuffer([]byte("invalid json")))
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, status)
	}
}
