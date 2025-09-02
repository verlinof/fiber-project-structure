package routes

import (
	"github.com/gofiber/fiber/v2"
	auth_http "github.com/verlinof/fiber-project-structure/internal/modules/auth/http"
	auth_http_route "github.com/verlinof/fiber-project-structure/internal/modules/auth/http/route"
	auth_service "github.com/verlinof/fiber-project-structure/internal/modules/auth/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

func InitRoute(app *fiber.App) {
	//EXAMPLE
	api := app.Group("/api")

	//Dependencies
	// redisManager := pkg_redis.NewRedisManager(redis_config.Config.Host, redis_config.Config.Password, redis_config.Config.Db)
	validator := pkg_validation.NewXValidator()

	// Services
	authService := auth_service.NewAuthService()

	// Auth
	authHandler := auth_http.NewAuthHandler(authService, validator)
	auth_http_route.AuthRoute(api, authHandler)
}
