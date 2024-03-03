package model

import (
	"gorm.io/gorm"
)

type Category struct {
	Name           string    `json:"name"`
	ThumbnailUrl   string    `json:"thumbnailUrl"`
	Slug           string    `json:"slug" gorm:"unique"`
	Description    string    `json:"description"`
	Products       []Product `json:"products"`
	OrderOfDisplay *uint     `json:"orderOfDisplay" gorm:"unique"`
}
