package categories

import (
	"gorm.io/gorm/clause"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(ctx *fiber.Ctx) error {
	var category model.Category

	if err := SelectCategoryByID(&category, ctx); err != nil {
		return helpers.SendNotFound(ctx, err)
	}

	// Set the order of display to nil, so it become available for another as there is a unique constraint
	category.OrderOfDisplay = nil

	if err := config.Database.Save(&category).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	categoryID := ctx.Params("categoryID")

	if err := config.Database.Clauses(clause.Returning{}).Delete(&category, categoryID).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	return sendCategory(ctx, category)
}
