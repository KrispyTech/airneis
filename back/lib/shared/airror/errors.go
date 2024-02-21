package airror

import (
	c "github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/pkg/errors"
)

func AutoMigrateError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.AutoMigrate], err)
}

func BootError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Boot], err)
}

func BootedError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Booted], err)
}

func BuildError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Build], err)
}

func FormatError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Format], err)
}

func GetError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Get], err)
}

func InitializeError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Initialize], err)
}

func LoadError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Load], err)
}

func LoadedError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Loaded], err)
}

func MarshallError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Marshall], err)
}

func MigrateError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Migrate], err)
}

func PrepareError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Prepare], err)
}

func ReadError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Read], err)
}

func StartedError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Started], err)
}

func UnmarshallError(parentFunctionName string, err error) error {
	return errors.Errorf("%s, %s %s", parentFunctionName, c.SelectTextUnableTo[c.Unmarshall], err)
}
