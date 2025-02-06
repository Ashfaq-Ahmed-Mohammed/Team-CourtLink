package DataBase

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SportSelection struct {
	Sport string `json:"sport"`
}

type CourtAvailability struct {
	CourtName   string `json:"CourtName"`
	CourtStatus uint   `json:"CourtStatus"`
}

type Customer struct {
	Customer_ID uint   `gorm:"column:Customer_ID;primaryKey;autoIncrement" json:"Customer_ID"`
	Name        string `gorm:"column:Name" json:"Name"`
	Contact     string `gorm:"column:Contact" json:"Contact"`
	Email       string `gorm:"column:Email" json:"Email"`
}

type Sport struct {
	Sport_ID          uint   `gorm:"column:Sport_ID;primaryKey;autoIncrement;unique;not null" json:"Sport_ID"`
	Sport_name        string `gorm:"column:Sport_name;unique;not null" json:"Sport_name"`
	Sport_Description string
}

type Court struct {
	Court_ID       uint   `gorm:"column:Court_ID;primaryKey;autoIncrement" json:"Court_ID"`
	Court_Name     string `gorm:"column:Court_Name;unique;not null" json:"Court_Name"`
	Court_Location string `gorm:"column:Court_Location;not null" json:"Court_Location"`
	Court_Capacity *int
	Court_Status   int    `gorm:"column:Court_Status;not null" json:"Court_Status"`
	Sport_id       *uint  `gorm:"column:Sport_id;index;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Sport_id"`
	Sport          *Sport `gorm:"foreignKey:Sport_ID"`
}

var DB *gorm.DB

func (Customer) TableName() string {
	return "Customer"
}

func (Sport) TableName() string {
	return "Sport"
}

func (Court) TableName() string {
	return "Court"
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("../CourtLink.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
