package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(ctx *fiber.Ctx) error {
	// We need to convert the page because query product must be an int
	page, errParsing := helpers.ConvertStringToInt(ctx.Query(constants.Page, constants.DefaultPageValue))
	if errParsing != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	products, err := queryProducts(page, ctx)

	if err != nil {
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return sendProducts(ctx, products)
}
