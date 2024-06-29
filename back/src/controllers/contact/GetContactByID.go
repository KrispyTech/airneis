package contact

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetContactByID(ctx *fiber.Ctx) error {
	var contact model.Contact

	if err := SelectContactTicketByID(&contact, ctx); err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return sendContactTicket(ctx, contact)
}
