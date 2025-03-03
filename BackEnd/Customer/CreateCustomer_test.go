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

	var err error
	DataBase.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	DataBase.DB.AutoMigrate(&DataBase.Customer{})
}

func TestCreateCustomer(t *testing.T) {
	setupTestDB()

	customerRequest := map[string]interface{}{
		"Name":    "Rohi B",
		"Email":   "rohb@example.com",
		"Contact": "1234567890",
	}

	body, _ := json.Marshal(customerRequest)
	req, err := http.NewRequest("POST", "/Customer", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateCustomer)
	handler.ServeHTTP(recorder, req)

	t.Logf("Response Body: %s", recorder.Body.String())

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, status)
	}

	var response map[string]interface{}
	json.Unmarshal(recorder.Body.Bytes(), &response)
	if response["message"] != "Customer record added successfully" {
		t.Errorf("unexpected response message: %v", response["message"])
	}

	var savedCustomer DataBase.Customer
	result := DataBase.DB.First(&savedCustomer, "email = ?", "rohb@example.com")

	t.Logf("Database Query Result: %+v", savedCustomer)

	if result.Error != nil {
		t.Errorf("Customer not found in database: %v", result.Error)
	} else if savedCustomer.Name != "Rohi B" || savedCustomer.Contact != "1234567890" {
		t.Errorf("Customer data mismatch: got %+v", savedCustomer)
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
