package main

import (
	"github.com/KrispyTech/airneis/controllers/categories"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/categories/", controllers.GetCategories)
	app.Get("/api/categories/", controllers.GetCategoryByID)
	app.Post("/api/categories/", controllers.CreateCategory)
	app.Patch("/api/categories/:categoryID", controllers.EditCategory)
	app.Delete("/api/categories/:categoryID", controllers.DeleteCategory)
}
