package pkg_validation

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func (v XValidator) Validate(data interface{}) error {
	var err error
	validationErrors := []ErrorResponse{}

	errs := v.validator.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag

			validationErrors = append(validationErrors, elem)
		}

		//Convert []ErrorResponse -> golang standard error
		errMsgs := make([]string, 0)

		for _, err := range validationErrors {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"%s Need to Implement '%v'",
				err.FailedField,
				err.Tag,
			))
		}
		// Separator
		errString := strings.Join(errMsgs, ";")
		err = errors.New(errString)
	}

	return err
}

// Adding custom validation for Struct
func (v XValidator) InitCustomValidation() {
	// Custom struct validation tag format (EXAMPLE)
	v.validator.RegisterValidation("example", func(fl validator.FieldLevel) bool {
		// Can be filled with conditional statement
		// // User.Age needs to fit our needs, 12-18 years old.
		// return fl.Field().Int() >= 12 && fl.Field().Int() <= 18
		return true
	})
}
