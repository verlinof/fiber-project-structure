package middleware

import (
	"fmt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"

	"github.com/verlinof/fiber-project-structure/configs/app_config"
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
		return c.Status(fiber.StatusUnauthorized).JSON(pkg_error.NewUnauthorized(fmt.Errorf("invalid JWT")))
	}
}
