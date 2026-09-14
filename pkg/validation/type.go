package pkg_validation

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type (
	XValidator struct {
		validator *validator.Validate
		db        *gorm.DB
	}

	ErrorResponse struct {
		FailedField string
		Tag         string
	}
)

// Initiate Validator
func NewXValidator(db *gorm.DB) XValidator {
	XValidator := XValidator{
		validator: validator.New(),
		db:        db,
	}

	XValidator.InitCustomValidation()

	return XValidator
}
