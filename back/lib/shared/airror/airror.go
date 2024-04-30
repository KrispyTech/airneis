package airror

import "fmt"

func New(err error, message string) error {
	if message != "" {
		return fmt.Errorf("%s: %s", message, err.Error())
	}

	return fmt.Errorf(err.Error())
}
