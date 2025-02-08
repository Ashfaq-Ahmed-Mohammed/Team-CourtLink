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
	Slots       []int  `json:"Slots"`
}

type Customer struct {
	Customer_ID uint   `gorm:"column:Customer_ID;primaryKey;autoIncrement" json:"Customer_ID"`
	Name        string `gorm:"column:Name" json:"name"`
	Contact     string `gorm:"column:Contact" json:"Contact"`
	Email       string `gorm:"column:Email" json:"email"`
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

type Court_TimeSlots struct {
	ID         uint   `gorm:"column:ID;primaryKey;autoIncrement" json:"ID"`
	Slot_8_9   int    `gorm:"column:slot_8_9;not null;default:1" json:"slot_8_9"`
	Slot_9_10  int    `gorm:"column:slot_9_10;not null;default:1" json:"slot_9_10"`
	Slot_10_11 int    `gorm:"column:slot_10_11;not null;default:1" json:"slot_10_11"`
	Slot_11_12 int    `gorm:"column:slot_11_12;not null;default:1" json:"slot_11_12"`
	Slot_12_13 int    `gorm:"column:slot_12_13;not null;default:1" json:"slot_12_13"`
	Slot_13_14 int    `gorm:"column:slot_13_14;not null;default:1" json:"slot_13_14"`
	Slot_14_15 int    `gorm:"column:slot_14_15;not null;default:1" json:"slot_14_15"`
	Slot_15_16 int    `gorm:"column:slot_15_16;not null;default:1" json:"slot_15_16"`
	Slot_16_17 int    `gorm:"column:slot_16_17;not null;default:1" json:"slot_16_17"`
	Slot_17_18 int    `gorm:"column:slot_17_18;not null;default:1" json:"slot_17_18"`
	Court_Name string `gorm:"column:Court_Name;unique;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"Court_Name"`
	Court      *Court `gorm:"foreignKey:Court_Name"`
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

func (Court_TimeSlots) TableName() string {
	return "Court_TimeSlots"
}

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("../CourtLink.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}
