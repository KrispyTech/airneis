package model

type ProductImage struct {
	ProductID uint    `json:"productID"`
	Product   Product `json:"product"`
	ImageURL  string  `json:"imageURL"`
}
