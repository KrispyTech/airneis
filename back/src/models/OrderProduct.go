package model

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	Order     Order
	ProductID uint
	Product   Product
	Quantity  int
}
