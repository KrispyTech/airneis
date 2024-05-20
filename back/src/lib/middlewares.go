package lib

import (
	"fmt"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (airneis *airneisFiberApp) setupLogger() fiber.Router {
	return airneis.app.Use(logger.New())
}

func (airneis *airneisFiberApp) setupCORS() fiber.Router {
	return airneis.app.Use(cors.New(
		cors.Config{
			AllowCredentials: true,
			AllowOriginsFunc: func(origin string) bool {
				log.Infof("Origin: %s", origin)

				return strings.Contains(origin, "http://localhost:3000")
			},
		}))
}

func authMiddleware(vc vault.VaultClient) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sessionToken, ok := helpers.GetSessionCookiesFromContext(ctx)
		if !ok {
			return helpers.SendUnauthorized(ctx, fmt.Errorf("unable to get session token"))
		}

		user, errToken := auth.SendSerializedUser(sessionToken, vc)
		if errToken != nil {
			return helpers.SendUnauthorized(ctx, errToken)
		}

		if !user.IsActive {
			return helpers.SendUnauthorized(ctx, fmt.Errorf("user is inactive"))
		}

		if err := helpers.AddUserToContext(ctx, user); err != nil {
			return err
		}

		return ctx.Next()
	}
}
