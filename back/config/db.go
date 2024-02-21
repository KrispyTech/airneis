package config

import (
	"fmt"
	"os"

	"github.com/KrispyTech/airneis/lib/shared/airror"
	c "github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	model "github.com/KrispyTech/airneis/src/db/models"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB

type NeonCreds struct {
	Username string
	Secret   string
	Server   string
}

func initializeDatabase(config Config, env string) (err error) {
	switch env {
	case c.StagingEnv:
		databaseURI := os.Getenv("DB_URI")
		if databaseURI == "" {
			return airror.GetError("os.Getenv", errors.Errorf("DB_URI missing"))
		}

		database, err = gorm.Open(postgres.Open(databaseURI))
		if err != nil {
			return airror.InitializeError("gorm.Open", err)
		}

	default:
		database, err = initializeNeonDatabase(config.Handler.VaultClient)
		if err != nil {
			return airror.WrapInitializeError("initializeNeonDatabase", err)
		}
	}

	log.Info(c.DatabaseLoaded)

	if err := database.AutoMigrate(&model.Product{},
		&model.Address{}, &model.Category{}, &model.Contact{},
		&model.Material{}, &model.Order{}, &model.Product{},
		&model.Status{}, &model.User{}); err != nil {
		return errors.Errorf("initDatabase, unable to auto migrate")
	}

	log.Info(c.AutomigrationsCompleted)

	return
}

func initializeNeonDatabase(vc vault.VaultClient) (*gorm.DB, error) {
	neonCred, err := GetNeonCreds(vc)
	if err != nil {
		return &gorm.DB{}, airror.WrapGetError("GetNeonCreds", err)
	}

	databaseURI := fmt.Sprintf("postgresql://%s:%s@%s?sslmode=require", neonCred.Username, neonCred.Secret, neonCred.Server)

	database, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, airror.LoadError("gorm.Open", err)
	}

	return database, nil
}

func GetNeonCreds(vc vault.VaultClient) (NeonCreds, error) {
	username, err := vc.ReadSecret("neon_username")
	if err != nil {
		return NeonCreds{}, airror.WrapGetError("vc.ReadSecret", err)
	}
	secret, err := vc.ReadSecret("neon_secret")
	if err != nil {
		return NeonCreds{}, airror.WrapGetError("vc.ReadSecret", err)
	}
	server, err := vc.ReadSecret("neon_server")
	if err != nil {
		return NeonCreds{}, airror.WrapGetError("vc.ReadSecret", err)
	}

	return NeonCreds{
		Username: username,
		Secret:   secret,
		Server:   server,
	}, nil
}
