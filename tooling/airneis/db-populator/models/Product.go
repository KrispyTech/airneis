package model

type Product struct {
	Name                    string         `json:"name"`
	Slug                    string         `json:"slug"`
	CategoryID              uint           `json:"categoryID"`
	Category                Category       `json:"category"`
	PriceWithoutTaxes       int            `json:"priceWithoutTaxes"`
	PriceWithTaxes          int            `json:"priceWithTaxes"`
	Description             string         `json:"description"`
	ThumbnailURL            string         `json:"thumbnailURL"`
	OutOfStock              bool         	 `json:"outOfStock"`
	Materials               []Material     `json:"materials"`
	HighlanderOrder 				uint           `json:"highlanderOrder"`
	OrderOfPriority         *uint          `json:"orderOfPriority"`
	ProductImages           []ProductImage `json:"productImages"`
}
