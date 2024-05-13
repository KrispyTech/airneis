package serializers

import (
	model "github.com/KrispyTech/airneis/src/models"
)

type SerializedUser struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	IsAdmin   bool   `json:"isAdmin"`
	IsActive  bool   `json:"isActive"`
}

func SerializeUser(user model.User) SerializedUser {
	return SerializedUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsActive:  user.IsActive,
	}
}
