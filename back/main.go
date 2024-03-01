package main

import (
	"github.com/KrispyTech/airneis/config"
	c "github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Backend configuration booting")
	_, err := config.InitializeConfig()
	if err != nil {
		log.Fatal("Unable to load config - ", err.Error())
	}

	app := fiber.New()
	log.Info("Backend application, fiber started")

	app.Get(c.HomeRoute, func(ctx *fiber.Ctx) error {
		return ctx.SendString(c.HelloWorld)
	})

	log.Info("Routes defined")
	log.Fatal(app.Listen(c.Port))
}
