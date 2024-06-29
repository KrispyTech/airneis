package users

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(ctx *fiber.Ctx) error {
	var user model.User

	if err := ctx.BodyParser(&user); err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	if err := checkUser(user); err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	user, err := fillUserInformation(ctx, user)
	if err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	if err := config.Database.Create(&user).Error; err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return nil
}
