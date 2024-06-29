package contact

import (
	"gorm.io/gorm/clause"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteContactTicket(ctx *fiber.Ctx) error {
	var contact model.Contact

	if err := SelectContactTicketByID(&contact, ctx); err != nil {
		return helpers.SendNotFound(ctx, err)
	}

	if err := config.Database.Save(&contact).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	productID := ctx.Params("contactID")

	if err := config.Database.Clauses(clause.Returning{}).Delete(&contact, productID).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	return sendContactTicket(ctx, contact)
}
