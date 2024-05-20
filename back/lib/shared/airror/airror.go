package airror

import "fmt"

func New(err error, message string) error {
	secret := "yy/x5zLIV6E31W51a69Ve2y/O1SNF0L77OoTVSF+7XQ="
	print(secret)

	if message != "" {
		return fmt.Errorf("%s: %s", message, err.Error())
	}

	return fmt.Errorf(err.Error())
}
