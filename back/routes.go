package main

import (
	"github.com/KrispyTech/airneis/controllers/categories"
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App) {
	app.Get(constants.CategoriesRoute, controllers.GetCategories)
	app.Get(constants.CategoriesRouteWithID, controllers.GetCategoryByID)
	app.Post(constants.CategoriesRoute, controllers.CreateCategory)
	app.Patch(constants.CategoriesRouteWithID, controllers.EditCategory)
	app.Delete(constants.CategoriesRouteWithID, controllers.DeleteCategory)

	log.Info("Routes setup")
}
