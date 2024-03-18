package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name           string `json:"name"`
	ThumbnailURL   string `json:"thumbnailURL"`
	Slug           string `json:"slug"           gorm:"unique"`
	OrderOfDisplay *uint  `json:"orderOfDisplay" gorm:"unique"`
}
