package Admin

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

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}
	db.AutoMigrate(&DataBase.Admin{})
	return db
}

func TestAdminLogin_Success(t *testing.T) {
	db := setupTestDB()
	DataBase.DB = db

	// Add test admin
	admin := DataBase.Admin{
		Username: "adminuser",
		Password: "adminpass",
	}
	db.Create(&admin)

	// Create request body
	loginReq := LoginRequest{
		Username: "adminuser",
		Password: "adminpass",
	}
	body, _ := json.Marshal(loginReq)

	req, err := http.NewRequest("POST", "/AdminLogin", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminLogin)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, status)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)
	if response["message"] != "Login successful" {
		t.Errorf("expected login success message, got: %v", response["message"])
	}
}

func TestAdminLogin_InvalidCredentials(t *testing.T) {
	db := setupTestDB()
	DataBase.DB = db

	// Add test admin
	admin := DataBase.Admin{
		Username: "adminuser",
		Password: "adminpass",
	}
	db.Create(&admin)

	// Wrong password
	loginReq := LoginRequest{
		Username: "adminuser",
		Password: "wrongpass",
	}
	body, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/AdminLogin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminLogin)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr.Code)
	}
}

func TestAdminLogin_InvalidUsername(t *testing.T) {
	db := setupTestDB()
	DataBase.DB = db

	// Valid admin in DB
	admin := DataBase.Admin{
		Username: "adminuser",
		Password: "adminpass",
	}
	db.Create(&admin)

	// Wrong username
	loginReq := LoginRequest{
		Username: "wronguser",
		Password: "adminpass",
	}
	body, _ := json.Marshal(loginReq)

	req, _ := http.NewRequest("POST", "/AdminLogin", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminLogin)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("expected status 401, got %d", rr.Code)
	}
}

func TestAdminLogin_InvalidJSON(t *testing.T) {
	req, _ := http.NewRequest("POST", "/AdminLogin", bytes.NewBuffer([]byte(`invalid-json`)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminLogin)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", rr.Code)
	}
}
