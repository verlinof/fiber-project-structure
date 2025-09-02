package user_route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verlinof/fiber-project-structure/internal/middleware"
	user_http "github.com/verlinof/fiber-project-structure/internal/modules/user/http"
)

func UserRoute(api fiber.Router, userHandler user_http.UserHandler) {
	userRoute := api.Group("/users")

	userRoute.Use(middleware.AuthMiddleware())
	userRoute.Get("/", userHandler.GetUsers)
	userRoute.Get("/:id", userHandler.GetUserbyID)
	userRoute.Post("/", middleware.RoleMiddleware("users.create"), userHandler.CreateUser)
	userRoute.Patch("/:id", middleware.RoleMiddleware("users.update"), userHandler.UpdateUser)
	userRoute.Patch("/change-password/:id", middleware.RoleMiddleware("users.update"), userHandler.ChangePassword)
	userRoute.Delete("/:id", middleware.RoleMiddleware("users.delete"), userHandler.DeleteUser)
}
