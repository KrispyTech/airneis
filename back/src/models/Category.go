package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name           string    `json:"name"`
	ThumbnailURL   string    `json:"thumbnailUrl"`
	Slug           string    `json:"slug"           gorm:"unique"`
	Description    string    `json:"description"`
	Products       []Product `json:"products"`
	OrderOfDisplay *uint     `json:"orderOfDisplay" gorm:"unique"`
}
