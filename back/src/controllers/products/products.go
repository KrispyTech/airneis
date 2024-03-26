package controllers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/pkg/config"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func sendProduct(ctx *fiber.Ctx, product model.Product) error {
	productJSON, err := json.Marshal(product)
	if err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(productJSON)
}

func sendProducts(ctx *fiber.Ctx, products []model.Product) error {
	jsonProducts, errMarshal := json.Marshal(products)
	if errMarshal != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(jsonProducts)
}

func SelectProductByID(product *model.Product, ctx *fiber.Ctx) error {
	productID := ctx.Params("productID")

	return config.Database.First(&product, productID).Error
}
