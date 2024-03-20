package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func (_v XValidator) validate(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem = ErrorResponse{
				FailedField: err.Field(),
				Tag:         err.Tag(),
				Value:       err.Value(),
				Error:       true,
			}

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

var Validator = &XValidator{
	validator: validate,
}

func ValidateInput(input interface{}) error {
	if errs := Validator.validate(input); len(errs) > 0 && errs[0].Error {
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
