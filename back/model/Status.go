package model

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name    string `json:"name"`
	IsOrder bool   `json:"isOrder" default:"false"`
}
