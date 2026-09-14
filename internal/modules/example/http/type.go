package example_http

import (
	example_service "github.com/verlinof/fiber-project-structure/internal/modules/example/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

type ExampleHandler struct {
	exampleService example_service.ExampleService
	xValidator     pkg_validation.XValidator
}

func NewHandler(exampleService example_service.ExampleService, xValidator pkg_validation.XValidator) ExampleHandler {
	return ExampleHandler{
		exampleService: exampleService,
		xValidator:     xValidator,
	}
}
