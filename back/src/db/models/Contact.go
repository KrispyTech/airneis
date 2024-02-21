package model

import (
	"gorm.io/gorm"
)

type Contact struct {
	gorm.Model
	SenderEmail    string `json:"senderEmail"`
	RecipientEmail string `json:"recipientEmail"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}
