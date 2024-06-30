package contact

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetContactTickets(ctx *fiber.Ctx) error {
	page, errParsing := helpers.ConvertStringToInt(ctx.Query(constants.Page, constants.DefaultPageValue))
	if errParsing != nil {
		return helpers.SendBadRequest(ctx, errParsing)
	}

	contactTickets, err := queryContactTickets(page, ctx)
	if err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return sendContactTickets(ctx, contactTickets)
}

func queryContactTickets(page int, ctx *fiber.Ctx) ([]model.Contact, error) {
	pagLimit := constants.PaginationLimit
	offset := (page - 1) * pagLimit

	contacts := make([]model.Contact, pagLimit)
	errQuery := config.Database.Limit(pagLimit).Offset(offset).Order("order_of_priority asc").
		Preload("Contact").Find(&contacts)
	if errQuery.Error != nil {
		return contacts, helpers.SendBadRequest(ctx, errQuery.Error)
	}

	return contacts, nil
}
