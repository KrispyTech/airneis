package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	"github.com/KrispyTech/airneis/pkg/validator"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if err := ctx.BodyParser(&product); err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	// We need to check if the category exist to link it to the product
	// It will also fill the category field of product that will also avoid to have validation issues with the category
	if err := config.Database.First(&product.Category, product.CategoryID).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.NotFoundStatus,
			constants.NotFoundMessage,
			ctx)
	}

	if err := validator.ValidateInput(product); err != nil {
		return helpers.SetStatusAndMessages(constants.BadRequestStatus, err.Error(), ctx)
	}

	if err := config.Database.Preload("category").Create(&product).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendProduct(ctx, product)
}
