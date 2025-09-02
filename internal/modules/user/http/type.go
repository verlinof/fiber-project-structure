package user_http

import (
	user_service "github.com/verlinof/fiber-project-structure/internal/modules/user/service"
	pkg_validation "github.com/verlinof/fiber-project-structure/pkg/validation"
)

type UserHandler struct {
	userService user_service.UserService
	xValidator  pkg_validation.XValidator
}

func NewUserHandler(userService user_service.UserService, xValidator pkg_validation.XValidator) UserHandler {
	return UserHandler{
		userService: userService,
		xValidator:  xValidator,
	}
}
