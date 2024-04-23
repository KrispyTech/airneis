package lib

import (
	"fmt"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/KrispyTech/airneis/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	log "github.com/sirupsen/logrus"
)

type airneisFiberApp struct {
	app      *fiber.App
	services Services
}

type Services struct {
	VaultClient vault.VaultClient
}

func InitializeFiberApp(services Services) error {
	airneis := airneisFiberApp{
		fiber.New(),
		services,
	}

	airneis.addLogger()
	airneis.setupRoutingConfig()
	airneis.app.GetRoutes()

	if err := airneis.app.Listen(constants.Port); err != nil {
		return fmt.Errorf("something went wrong when initializing Fiber App. %s", err.Error())
	}

	return nil
}

func (airneis *airneisFiberApp) setupRoutingConfig() fiber.Router {
	router := airneis.app.Group("/api")
	log.Info("All routes has been grouped to /api")

	return SetupRoutes(router)
}

func SetupRoutes(router fiber.Router) fiber.Router {
	routes.SetProductsRoutes(router)
	routes.SetProductsRoutes(router)
	routes.SetUsersRoutes(router)
	routes.SetSessionsRoutes(router)
	log.Info("Routes have been set")

	return router
}

func (airneis *airneisFiberApp) addLogger() fiber.Router {
	return airneis.app.Use(logger.New())
}
