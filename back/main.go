package main

import (
	"github.com/KrispyTech/airneis/pkg/config"
	"github.com/KrispyTech/airneis/src/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Backend configuration booting")
	conf, err := config.InitializeConfig()
	if err != nil {
		log.Fatal("Unable to load config - ", err.Error())
	}

	externalServices := lib.Services{
		VaultClient: conf.Handler.VaultClient,
		SendgridClient: conf.Handler.SendgridClient,
	}

	if err := lib.InitializeFiberApp(externalServices); err != nil {
		log.Fatalf("Unable to create fiber app, %s", err.Error())
	}
}
