package example_route

import (
	example_http "github.com/verlinof/fiber-project-structure/internal/modules/example/http"
	"github.com/gofiber/fiber/v2"
)

// InitRoute registers routes for the example module
func InitRoute(router fiber.Router, exampleHandler example_http.ExampleHandler) {
	examples := router.Group("/examples")

	examples.Get("/", exampleHandler.GetAll)
	examples.Get("/:id", exampleHandler.GetByID)
	examples.Post("/", exampleHandler.Create)
	examples.Put("/:id", exampleHandler.Update)
	examples.Delete("/:id", exampleHandler.Delete)
}
