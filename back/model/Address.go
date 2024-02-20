package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID      uint    `json:"userID"`
	User        User    `json:"user"`
	Address     string  `json:"address"`
	Address2    *string `json:"address2"`
	City        string  `json:"city"`
	ZipCode     string  `json:"zipCode"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	PhoneNumber string  `json:"phoneNumber"`
	Comment     *string `json:"comment"`
	IsActive    bool    `json:"isActive" gorm:"default:true"`
}
