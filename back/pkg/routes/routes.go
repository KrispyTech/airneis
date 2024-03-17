package routes

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func SetRoutes(app *fiber.App) {
	app.Group("/api")
	setCategoriesRoutes(app)
	log.Info("Routes have been set")
}
