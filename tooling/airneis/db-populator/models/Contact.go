package model

type Contact struct {
	SenderEmail    string `json:"senderEmail"`
	RecipientEmail string `json:"recipientEmail"`
	Subject        string `json:"subject"`
	Message        string `json:"message"`
}
