package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	IsHighlanderOfTheMoment bool           `json:"isHighlanderOfTheMoment"                                  gorm:"default:false"`
	Name                    string         `json:"name"                    validate:"required,min=2,max=50"`
	CategoryID              uint           `json:"categoryID"              validate:"required,min=1"`
	Category                Category       `json:"category"`
	PriceWithoutTaxes       int            `json:"PriceWithoutTaxes"       validate:"required,min=1"`
	PriceWithTaxes          int            `json:"priceWithTaxes"          validate:"required,min=1"`
	Description             string         `json:"description"`
	ThumbnailURL            string         `json:"thumbnailURL"            validate:"required"`
	OutOfStock              string         `json:"outOfStock"                                               gorm:"default:false"`
	Materials               []Material     `json:"materials"                                                gorm:"many2many:materials_of_products"`
	OrderOfPriority         *uint          `json:"orderOfPriority"         validate:"required"`
	ProductImages           []ProductImage `json:"productImages"                                            gorm:"many2many:images_of_products"`
}
