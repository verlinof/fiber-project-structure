package auth_http

import (
	auth_service "github.com/verlinof/fiber-project-structure/internal/modules/auth/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

type AuthHandler struct {
	authService auth_service.AuthService
	xValidator  pkg_validation.XValidator
}

func NewAuthHandler(authService auth_service.AuthService, xValidator pkg_validation.XValidator) AuthHandler {
	return AuthHandler{
		authService: authService,
		xValidator:  xValidator,
	}
}
