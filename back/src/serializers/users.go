package serializers

import (
	model "github.com/KrispyTech/airneis/src/models"
)

type SerializedUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"     gorm:"unique"`
	IsAdmin   bool   `json:"isAdmin"   gorm:"default:false"`
	IsActive  bool   `json:"isActive"  gorm:"default:false"`
}

func SerializeUser(user model.User) SerializedUser {
	return SerializedUser{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		IsAdmin:   user.IsAdmin,
		IsActive:  user.IsActive,
	}
}
