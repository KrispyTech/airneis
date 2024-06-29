package contact

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/pkg/config"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func sendContactTicket(ctx *fiber.Ctx, contact model.Contact) error {
	contactJSON, err := json.Marshal(contact)
	if err != nil {
		return helpers.SendInternalError(ctx, err)
	}

	return ctx.Send(contactJSON)
}

func sendContactTickets(ctx *fiber.Ctx, contacts []model.Contact) error {
	jsonContacts, errMarshal := json.Marshal(contacts)
	if errMarshal != nil {
		return helpers.SendInternalError(ctx, errMarshal)
	}

	return ctx.Send(jsonContacts)
}

func SelectContactTicketByID(contact *model.Contact, ctx *fiber.Ctx) error {
	productID := ctx.Params("contactID")

	return config.Database.Preload("Contact").First(&contact, productID).Error
}
