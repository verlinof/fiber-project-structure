package auth_http_route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/internal/middleware"
	auth_http "github.com/verlinof/fiber-project-structure/internal/modules/auth/http"
)

func AuthRoute(router fiber.Router, authHandler auth_http.AuthHandler) {
	authRoute := router.Group("/")

	authRoute.Post("/login", authHandler.Login)

	//Contoh penggunaan Middleware by Permission Name
	authRoute.Use(middleware.AuthMiddleware())
	authRoute.Get("/tes", middleware.RoleMiddleware("users.create"), authHandler.Tes) // Tag adalah nama permission
}
