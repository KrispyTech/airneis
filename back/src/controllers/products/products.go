package controllers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func sendProduct(ctx *fiber.Ctx, product model.Product) error {
	categoryJSON, err := json.Marshal(product)
	if err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(categoryJSON)
}

func sendProducts(ctx *fiber.Ctx, products []model.Product) error {
	jsonCategories, errMarshal := json.Marshal(products)
	if errMarshal != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(jsonCategories)
}

func selectProductByID(product *model.Product, ctx *fiber.Ctx) error {
	categoryID := ctx.Params("productID")

	return config.Database.Preload("Category").First(&product, categoryID).Error
}

func queryProducts(page int, ctx *fiber.Ctx) ([]model.Product, error) {
	pagLimit := constants.PaginationLimit
	offset := (page - 1) * pagLimit

	products := make([]model.Product, pagLimit)
	query := config.Database.
		Preload("Category").Limit(pagLimit).Offset(offset).Order("order_of_priority asc").Find(&products)
	if query.Error != nil {
		ctx.Status(constants.BadRequestStatus)

		return products, ctx.SendString(constants.BadRequestMessage)
	}

	return products, nil
}
