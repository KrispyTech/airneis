package config

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDatabase() (db *gorm.DB, err error) {
	db, err = gorm.Open(postgres.Open(os.Getenv("DB_URI")), &gorm.Config{})
	if err != nil {
		return &gorm.DB{}, errors.Wrapf(err, "InitDatabase, unable to load database")
	}

	logrus.Info("Database has been loaded")

	return db, nil
}
