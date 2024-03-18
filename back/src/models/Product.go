package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name                    string         `json:"name"`
	Slug                    string         `json:"slug"                    gorm:"unique"`
	CategoryID              uint           `json:"categoryID"`
	Category                Category       `json:"category"`
	PriceWithoutTaxes       float32        `json:"PriceWithoutTaxes"`
	PriceWithTaxes          float32        `json:"priceWithTaxes"`
	Description             string         `json:"description"`
	ThumbnailURL            string         `json:"thumbnailURL"`
	OutOfStock              string         `json:"outOfStock"              gorm:"default:false"`
	Materials               []Material     `json:"materials"               gorm:"many2many:materials_of_products"`
	IsHighlanderOfTheMoment bool           `json:"isHighlanderOfTheMoment" gorm:"default:false"`
	OrderOfPriority         *uint          `json:"orderOfPriority"`
	ProductImages           []ProductImage `json:"productImages"           gorm:"many2many:images_of_products"`
}
