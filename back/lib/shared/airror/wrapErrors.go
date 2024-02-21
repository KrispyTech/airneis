package airror

import (
	c "github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/pkg/errors"
)

func WrapAutoMigrateError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.AutoMigrate])
}

func WrapBootError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Boot])
}

func WrapBootedError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Booted])
}

func WrapBuildError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Build])
}

func WrapFormatError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Format])
}

func WrapGetError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Get])
}

func WrapInitializeError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Initialize])
}

func WrapLoadError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Load])
}

func WrapLoadedError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Loaded])
}

func WrapMarshallError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Marshall])
}

func WrapMigrateError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Migrate])
}

func WrapPrepareError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Prepare])
}

func WrapReadError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectUnableToUpon[c.Read])
}

func WrapStartError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectTextUnableTo[c.Started])
}

func WrapUnmarshallError(parentFunctionName string, err error) error {
	return errors.Wrapf(err, "%s, %s", parentFunctionName, c.SelectTextUnableTo[c.Unmarshall])
}
