package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetProductByID(ctx *fiber.Ctx) error {
	var product model.Product

	if err := SelectProductByID(&product, ctx); err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendProduct(ctx, product)
}
