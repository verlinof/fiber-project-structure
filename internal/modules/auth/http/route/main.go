package auth_route

import (
	"github.com/verlinof/fiber-project-structure/db"
	"github.com/verlinof/fiber-project-structure/internal/middleware"
	auth_http "github.com/verlinof/fiber-project-structure/internal/modules/auth/http"
	"github.com/gofiber/fiber/v2"
)

func InitRoute(router fiber.Router, authHandler auth_http.AuthHandler) {
	router.Post("/register", authHandler.Register)
	router.Post("/login", authHandler.Login)
	router.Get("/profile", middleware.AuthMiddleware(), middleware.CheckUserExists(db.GetDB()), authHandler.GetProfile)
}
