package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name           string    `json:"name"`
	ThumbnailUrl   string    `json:"thumbnailUrl"`
	Slug           string    `json:"slug" gorm:"unique"`
	Products       []Product `json:"products"`
	OrderOfDisplay *uint     `json:"orderOfDisplay" gorm:"unique"`
}
