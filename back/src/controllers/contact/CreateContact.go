package contact

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateContactTicket(ctx *fiber.Ctx) error {
	var contact model.Contact

	if err := ctx.BodyParser(&contact); err != nil {
		return helpers.SendBadRequest(ctx, err)
	}
	

	if err := config.Database.Create(&contact).Error; err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return sendContactTicket(ctx, contact)
}
