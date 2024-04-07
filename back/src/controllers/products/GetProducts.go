package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetProducts(ctx *fiber.Ctx) error {
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

func queryProducts(page int, ctx *fiber.Ctx) ([]model.Product, error) {
	pagLimit := constants.PaginationLimit
	offset := (page - 1) * pagLimit

	products := make([]model.Product, pagLimit)
	errQuery := config.Database.Limit(pagLimit).Offset(offset).Order("order_of_priority asc").Find(&products)
	if errQuery.Error != nil {
		ctx.Status(constants.BadRequestStatus)

		return products, ctx.SendString(constants.BadRequestMessage)
	}

	return products, nil
}
