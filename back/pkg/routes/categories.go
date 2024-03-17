package routes

import (
	controllers "github.com/KrispyTech/airneis/src/controllers/categories"
	"github.com/gofiber/fiber/v2"
)

const (
	categoriesRoute       = "categories"
	categoriesRouteWithID = "categories/:categoryID"
)

func setCategoriesRoutes(app fiber.Router) {
	app.Get(categoriesRoute, controllers.GetCategories)
	app.Get(categoriesRouteWithID, controllers.GetCategoryByID)
	app.Post(categoriesRoute, controllers.CreateCategory)
	app.Patch(categoriesRouteWithID, controllers.EditCategory)
	app.Delete(categoriesRouteWithID, controllers.DeleteCategory)
}
