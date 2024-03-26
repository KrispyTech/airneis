package model

import (
	"time"
)

type Order struct {
	UserID            uint       `json:"userID"`
	User              User       `json:"user"`
	Products          []Product  `json:"products"`
	StatusID          uint       `json:"statusID"`
	Status            Status     `json:"status"`
	ShippingAddressID uint       `json:"shippingAddressID"`
	ShippingAddress   *Address   `json:"shippingAddress"`
	BillingAddressID  uint       `json:"billingAddressID"`
	BillingAddress    *Address   `json:"billingAddress"`
	FinalizedAt       *time.Time `json:"finalizedAt"`
}
