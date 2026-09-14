package pkg_validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (v XValidator) Validate(data interface{}) error {
	validationErrors := v.validator.Struct(data)
	if validationErrors != nil {

		// Cast the error to validator.ValidationErrors
		errs := validationErrors.(validator.ValidationErrors)

		// Create a slice to hold the custom error messages
		var errorMessages []string

		for _, err := range errs {
			// Use a switch statement to create a custom message for each validation tag
			var message string
			switch err.Tag() {
			case "email":
				message = fmt.Sprintf("%s must be a valid email address.", err.Field())
			default:
				// A default message for any other validation errors
				message = fmt.Sprintf("'%s' need '%s' tag.", err.Field(), err.Tag())
			}
			errorMessages = append(errorMessages, message)
		}

		// Join all error messages into a single string
		return errors.New(strings.Join(errorMessages, "; "))
	}

	// Return nil if there are no validation errors
	return nil
}
