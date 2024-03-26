package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name                    string         `json:"name"`
	Slug                    string         `json:"slug"                    gorm:"unique"`
	CategoryID              uint           `json:"categoryID"`
	Category                Category       `json:"category"`
	PriceWithoutTaxes       int            `json:"PriceWithoutTaxes"`
	PriceWithTaxes          int            `json:"priceWithTaxes"`
	Description             string         `json:"description"`
	ThumbnailURL            string         `json:"thumbnailURL"`
	OutOfStock              bool         	 `json:"outOfStock"              gorm:"default:false"`
	Materials               []Material     `json:"materials"               gorm:"many2many:materials_of_products"`
	HighlanderOrder 				uint           `json:"highlanderOrder"         gorm:"default:0"`
	OrderOfPriority         *uint          `json:"orderOfPriority"`
	ProductImages           []ProductImage `json:"productImages"           gorm:"many2many:images_of_products"`
}
