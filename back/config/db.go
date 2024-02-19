package config

import (
	"fmt"
	"os"

	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/KrispyTech/airneis/model"
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

func GetNeonCreds(vc vault.VaultClient) (NeonCreds, error) {
	username, err := vc.ReadSecret("neon_username")
	if err != nil {
		return NeonCreds{}, errors.Wrapf(err, "GetNeonCreds, unable to get username")
	}
	secret, err := vc.ReadSecret("neon_secret")
	if err != nil {
		return NeonCreds{}, errors.Wrapf(err, "GetNeonCreds, unable to get secret")
	}
	server, err := vc.ReadSecret("neon_server")
	if err != nil {
		return NeonCreds{}, errors.Wrapf(err, "GetNeonCreds, unable to get server")
	}

	return NeonCreds{
		Username: username,
		Secret:   secret,
		Server:   server,
	}, nil
}

func InitNeonDatabase(vc vault.VaultClient) (*gorm.DB, error) {
	neonCred, err := GetNeonCreds(vc)
	if err != nil {
		return &gorm.DB{}, errors.Wrapf(err, "InitDB, unable to get credentials")
	}
	databaseURI := fmt.Sprintf("postgresql://%s:%s@%s?sslmode=require", neonCred.Username, neonCred.Secret, neonCred.Server)

	database, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, errors.Wrapf(err, "InitDatabase, unable to load database")
	}

	return database, nil
}

func InitDatabase(config Config, env string) (err error) {
	switch env {
	case stagingEnv:
		databaseURI := os.Getenv("DB_URI")
		if databaseURI == "" {
			return errors.Wrapf(err, "InitDatabase, unable to get database URI")
		}

		database, err = gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
		if err != nil {
			return errors.Wrapf(err, "InitDatabase, unable to load database")
		}

	default:
		database, err = InitNeonDatabase(config.Handler.VaultClient)
		if err != nil {
			return errors.Wrapf(err, "InitDatabase, unable to load database")
		}

	}

	log.Info("Database has been loaded")

	if err = database.AutoMigrate(); err != nil {
		return errors.Wrapf(err, "InitDatabase, unable to do auto migration")
	}

	return
}

func AutoMigrate(database *gorm.DB) (err error) {
	if err := database.AutoMigrate(
		&model.Address{},
		&model.Category{},
		&model.Contact{},
		&model.Material{},
		&model.Order{},
		&model.Product{},
		&model.Status{},
		&model.User{},
	); err != nil {
		return err
	}

	return
}
