package model

import (
	"gorm.io/gorm"
	"time"
)

type Order struct {
	gorm.Model
	UserID            uint       `json:"userID"`
	User              User       `json:"user"`
	Products          []Product  `json:"products" gorm:"many2many:products_of_orders"`
	StatusID          uint       `json:"statusID"`
	Status            Status     `json:"status"`
	ShippingAddressID uint       `json:"shippingAddressID"`
	ShippingAddress   *Address   `json:"shippingAddress"`
	BillingAddressID  uint       `json:"billingAddressID"`
	BillingAddress    *Address   `json:"billingAddress"`
	FinalizedAt       *time.Time `json:"finalizedAt"`
}
