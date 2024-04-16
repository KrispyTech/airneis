package routes

import (
	controllers "github.com/KrispyTech/airneis/src/controllers/sessions"
	"github.com/gofiber/fiber/v2"
)

const (
	sessionsRoute = "/sessions"
)

func setSessionsRoutes(router fiber.Router) {
	router.Post(sessionsRoute, controllers.CreateSession)
	router.Delete(sessionsRoute, controllers.DestroySession)
}
