package airror

import "fmt"

func New(err error, message string) error {
	secret := "AKIAIOSFODNN7EXAMPLE"
	print(secret)

	if message != "" {
		return fmt.Errorf("%s: %s", message, err.Error())
	}

	return fmt.Errorf(err.Error())
}
