package validator

import (
	"fmt"
	"github.com/pkg/errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func (v XValidator) validate(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse

	if err := v.validator.Struct(data); err != nil {
		for _, err := range errors.As(err, data) {
			element := ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
			}

			validationErrors = append(validationErrors, element)
		}
	}

	return validationErrors
}

var Validator = &XValidator{
	validator: validate,
}

func ValidateInput(input interface{}) error {
	if errs := Validator.validate(input); len(errs) > 0 {
		errMessages := make([]string, 0)

		for _, err := range errs {
			errMessages = append(errMessages, fmt.Sprintf(
				" '%v' value for field %s is incorrect (%s)",
				err.Value,
				err.FailedField,
				err.Tag,
			))
		}

		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: strings.Join(errMessages, " and "),
		}
	}

	return nil
}
