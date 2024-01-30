package main

import (
	"airneis/config"

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

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	log.Info("Routes defined")
	log.Fatal(app.Listen(":3000"))
}
