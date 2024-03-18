package routes

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetRoutes(app *fiber.App) {
	apiRouter := app.Group("/api")
	setCategoriesRoutes(apiRouter)
	SetupProductsRoutes(apiRouter)
	log.Info("Routes have been set")
}
