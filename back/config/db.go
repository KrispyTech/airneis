package config

import (
	"os"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(config Config) (Config, error) {
	var err error

	databaseURI := os.Getenv("DB_URI")
	if databaseURI == "" {
		return Config{}, errors.Wrapf(err, "InitDatabase, unable to get database URI")
	}

	db, err := gorm.Open(postgres.Open(databaseURI), &gorm.Config{})
	if err != nil {
		return Config{}, errors.Wrapf(err, "InitDatabase, unable to load database")
	}

	log.Info("Database has been loaded")
	config.DB = db

	return config, nil
}
