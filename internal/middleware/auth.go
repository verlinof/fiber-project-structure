package middleware

import (
	"errors"
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"

	"github.com/verlinof/fiber-project-structure/configs/app_config"
	auth_model "github.com/verlinof/fiber-project-structure/internal/modules/auth/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
)

// AuthMiddleware authenticates the JWT token
func AuthMiddleware() fiber.Handler {
	return jwtware.New(jwtware.Config{
		ContextKey:   "token",
		SigningKey:   jwtware.SigningKey{Key: []byte(app_config.Config.JwtSecretKey)},
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	} else {
		c.Status(fiber.StatusUnauthorized)
		return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("Unauthorized")))
	}
}

func CheckUserExists(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get token from context, it's set by the jwtware middleware
		token := c.Locals("token").(*jwt.Token)

		// Get claims from token
		claims := token.Claims.(jwt.MapClaims)

		// Get user ID from claims. JWT numbers are float64 by default.
		userIDFloat, ok := claims["id_user"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("invalid token claims")))
		}
		userID := int(userIDFloat)

		// Check if user exists in the database
		var user auth_model.User
		if err := db.Table("users").First(&user, "id = ?", userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("unauthorized")))
			}
			// Any other database error
			return c.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
		}

		// Proceed to the next handler
		return c.Next()
	}
}
