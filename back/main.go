package main

import (
	"github.com/KrispyTech/airneis/config"
	neon "github.com/KrispyTech/airneis/db"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Backend configuration booting")
	config, err := config.InitializeConfig()
	if err != nil {
		log.Fatal("Unable to load config - ", err.Error())
	}

	app := fiber.New()
	log.Info("Backend application, fiber started")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	db, err := neon.InitDB(config.Handler.VaultClient)
	if err != nil {
		log.Fatal("Unable to connect to DB", err.Error())
	}

	if err := config.DB.AutoMigrate( /* PLACE HERE THE STRUCTS OF THE MODEL*/ ); err != nil {
		log.Fatal("Unable to run automigrate", err.Error())
	}

	log.Info(neon.CheckVersion(db))

	log.Info("Routes defined")
	log.Fatal(app.Listen(":3000"))
}
