package main

import (
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	app := fiber.New()

	log.Info("Backend configuration booting")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Hi bshsssdsa bjdsbbdassskj")
	})

	log.Fatal(app.Listen(":3000"))
}
