package model

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name string `json:"name"`
}
