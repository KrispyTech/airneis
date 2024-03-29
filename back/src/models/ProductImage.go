package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductID uint    `json:"productID"`
	Product   Product `json:"product"`
	ImageURL  string  `json:"imageURL"`
}
