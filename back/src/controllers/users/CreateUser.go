package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var (
		user model.User
	)

	if err := ctx.BodyParser(&user); err != nil {
		return helpers.SendBadRequest(ctx)
	}

	hashedPassword, salt, errorGetPassword := helpers.GetHashedPasswordAndSaltFromContext(ctx)

	if errorGetPassword != nil {
		return helpers.SendBadRequest(ctx)
	}

	user.PasswordHash, user.PasswordSalt = hashedPassword, salt

	if err := config.Database.Create(&user).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return nil
}
