package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	"github.com/KrispyTech/airneis/pkg/validator"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if err := selectProductByID(&product, ctx); err != nil {
		return helpers.SetStatusAndMessages(
			constants.NotFoundStatus,
			constants.NotFoundMessage,
			ctx)
	}

	if err := ctx.BodyParser(&product); err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	if err := validator.ValidateInput(product); err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			err.Error(),
			ctx)
	}

	if err := config.Database.Save(&product).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendProduct(ctx, product)
}
