package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name                    string         `json:"name"`
	Slug                    string         `json:"slug" gorm:"unique"`
	CategoryID              uint           `json:"categoryID"`
	Category                Category       `json:"category"`
	PriceWithoutTaxes       int            `json:"PriceWithoutTaxes"`
	PriceWithTaxes          int            `json:"priceWithTaxes"`
	Description             string         `json:"description"`
	ThumbnailUrl            string         `json:"thumbnailUrl"`
	OutOfStock              string         `json:"outOfStock" gorm:"default:false"`
	Materials               []Material     `json:"materials" gorm:"many2many:materials_of_products"`
	IsHighlanderOfTheMoment bool           `json:"isHighlanderOfTheMoment" gorm:"default:false"`
	OrderOfPriority         int            `json:"orderOfPriority"`
	ProductImages           []ProductImage `json:"productImages" gorm:"many2many:images_of_products"`
}