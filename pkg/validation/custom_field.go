package pkg_validation

import (
	"errors"

	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Adding custom validation for Struct
func (v XValidator) InitCustomValidation() {
	v.validator.RegisterValidation("user_exist", func(fl validator.FieldLevel) bool {
		// Query to Database
		var user auth_model.UserResponse

		if err := v.db.Where("id = ?", fl.Field().Int()).First(&user).Error; err != nil {
			return !errors.Is(err, gorm.ErrRecordNotFound)
		}

		return true
	})
}
