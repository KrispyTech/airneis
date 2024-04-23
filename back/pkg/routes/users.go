package routes

import (
	controllers "github.com/KrispyTech/airneis/src/controllers/users"
	"github.com/gofiber/fiber/v2"
)

const (
	userRoute = "users"
)

func SetUsersRoutes(router fiber.Router) {
	router.Post(userRoute, controllers.CreateUser)
}
