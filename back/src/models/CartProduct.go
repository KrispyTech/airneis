package model

import "gorm.io/gorm"

type OrderProduct struct {
	gorm.Model
	OrderID   uint
	ProductID uint
	quantity  int
}
