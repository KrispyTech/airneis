package model

type User struct {
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	PasswordHash string  `json:"passwordHash"`
	PasswordSalt string  `json:"passwordSalt"`
	IsAdmin      bool    `json:"isAdmin"`
	IsActive     bool    `json:"isActive"`
	Orders       []Order `json:"orders"`
}
