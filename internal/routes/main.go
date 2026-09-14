package routes

import (
	"github.com/verlinof/fiber-project-structure/db"
	auth_http "github.com/verlinof/fiber-project-structure/internal/modules/auth/http"
	auth_route "github.com/verlinof/fiber-project-structure/internal/modules/auth/http/route"
	auth_service "github.com/verlinof/fiber-project-structure/internal/modules/auth/service"
	example_http "github.com/verlinof/fiber-project-structure/internal/modules/example/http"
	example_route "github.com/verlinof/fiber-project-structure/internal/modules/example/http/route"
	example_service "github.com/verlinof/fiber-project-structure/internal/modules/example/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(app *fiber.App) {
	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "healthy",
			"service": "api-service",
		})
	})

	v1 := app.Group("/v1")

	// Global Dependencies
	validator := pkg_validation.NewXValidator(db.GetDB())

	// Services
	authService := auth_service.NewService(db.GetDB())
	exampleService := example_service.NewService(db.GetDB())

	// Auth Module Routes
	authHandler := auth_http.NewHandler(authService, validator)
	auth_route.InitRoute(v1, authHandler)

	// Example Module Routes (Contoh pendaftaran modul baru)
	exampleHandler := example_http.NewHandler(exampleService, validator)
	example_route.InitRoute(v1, exampleHandler)
}
