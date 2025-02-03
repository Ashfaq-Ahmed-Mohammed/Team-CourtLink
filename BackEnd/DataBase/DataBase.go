package DataBase

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Customer struct {
	Customer_ID uint   `gorm:"column:Customer_ID;primaryKey;autoIncrement" json:"Customer_ID"`
	Name        string `gorm:"column:Name" json:"Name"`
	Contact     string `gorm:"column:Contact" json:"Contact"`
	Email       string `gorm:"column:Email" json:"Email"`
}

var DB *gorm.DB

func (Customer) TableName() string {
	return "Customer"
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("../CourtLink.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
