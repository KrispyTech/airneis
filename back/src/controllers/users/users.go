package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func fillUserInformation(ctx *fiber.Ctx, user model.User) (model.User, error) {
	authData, err := helpers.GetAuthContext(ctx)
	if err != nil {
		return model.User{}, helpers.SendInternalError(ctx)
	}

	hashedPassword, salt, err := auth.CreatePassword(authData)
	if err != nil {
		return model.User{}, helpers.SendBadRequest(ctx)
	}

	user.PasswordHash, user.PasswordSalt = hashedPassword, salt

	return user, nil
}

func GetUserByEmailOrError(email string) (model.User, error) {
	var user model.User

	if err := config.Database.Where("email = ?", email).First(&user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
