package middlewares

import (
	"fmt"
	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Middleware(app *fiber.App) {
	app.Use(
		cors.New(),
	)
}

func AuthMiddleware(vc vault.VaultClient) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		fmt.Println(ctx.Cookies("session"))
		user, errToken := auth.CheckToken(ctx.Cookies("session"), vc)
		if errToken != nil {
			return helpers.SendUnauthorized(ctx, errToken)
		}

		if !user.IsActive {
			return helpers.SendUnauthorized(ctx, fmt.Errorf("User is not active"))
		}

		ctx.Locals("user", user)

		return ctx.Next()
	}
}
