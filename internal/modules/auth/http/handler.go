package auth_http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (authHandler AuthHandler) Tes(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("tes")
}

func (authHandler AuthHandler) Login(ctx *fiber.Ctx) error {
	var authRequest auth_model.LoginRequest
	err := ctx.BodyParser(&authRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Validate
	err = authHandler.xValidator.Validate(authRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// Call Service
	authResponse, err := authHandler.authService.Login(ctx.Context(), authRequest)
	if err != nil {
		if err == gorm.ErrRecordNotFound || err == bcrypt.ErrMismatchedHashAndPassword {
			return ctx.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("username or password is incorrect")))
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(authResponse)
}
