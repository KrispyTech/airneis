package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func FiberMiddleware(app *fiber.App) {
	app.Use(
		cors.New(),
	)
}
