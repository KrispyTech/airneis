package contact

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditContactTicket(ctx *fiber.Ctx) error {
	var contact model.Contact

	if errGet := SelectContactTicketByID(&contact, ctx); errGet != nil {
		return helpers.SendNotFound(ctx, errGet)
	}

	if errParsing := ctx.BodyParser(&contact); errParsing != nil {
		return helpers.SendBadRequest(ctx, errParsing)
	}

	if err := config.Database.Save(&contact).Error; err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return sendContactTicket(ctx, contact)
}
