package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name           string `json:"name"           validate:"required,min=2,max=50"`
	ThumbnailURL   string `json:"thumbnailURL"   validate:"required,min=2"`
	OrderOfDisplay *uint  `json:"orderOfDisplay" validate:"required"              gorm:"unique"`
}
