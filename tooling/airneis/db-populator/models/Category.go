package model

type Category struct {
	Name           string `json:"name"`
	ThumbnailURL   string `json:"thumbnailURL"`
	Slug           string `json:"slug"`
	OrderOfDisplay *uint  `json:"orderOfDisplay"`
}
