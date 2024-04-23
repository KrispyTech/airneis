package routes

import (
	controllers "github.com/KrispyTech/airneis/src/controllers/sessions"
	"github.com/gofiber/fiber/v2"
)

const (
	sessionsRoute = "sessions"
)

func SetSessionsRoutes(router fiber.Router) fiber.Router {
	router.Post(sessionsRoute, controllers.CreateSession)
	router.Delete(sessionsRoute, controllers.DestroySession)

	return router
}
