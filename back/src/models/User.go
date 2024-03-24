package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"        gorm:"unique"`
	PasswordHash string  `json:"passwordHash"`
	PasswordSalt string  `json:"passwordSalt"`
	IsAdmin      bool    `json:"isAdmin"      gorm:"default:false"`
	IsActive     bool    `json:"isActive"     gorm:"default:false"`
	Orders       []Order `json:"orders"       gorm:"many2many:order_of_user"`
}
