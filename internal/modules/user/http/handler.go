package user_http

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	user_model "github.com/verlinof/fiber-project-structure/internal/modules/user/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	"gorm.io/gorm"
)

func (h UserHandler) GetUsers(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per_page", 10)
	idPuskesmas := ctx.QueryInt("id_puskesmas", 0)

	var users *pkg_success.PaginationData
	var err error

	if idPuskesmas != 0 {
		users, err = h.userService.GetUserbyPuskesmas(ctx.Context(), idPuskesmas, page, perPage)
	} else {
		users, err = h.userService.GetAllUsers(ctx.Context(), page, perPage)
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(users)
}

func (h UserHandler) GetUserbyID(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id")))
	}

	// Call Service
	userResponse, err := h.userService.GetUserbyID(ctx.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_success.SuccessGetData(userResponse))
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var req user_model.CreateUserRequest

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Validate
	err = h.xValidator.Validate(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// Call Service
	userResponse, err := h.userService.CreateUser(ctx.Context(), req)
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("username already exists")))
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	return ctx.Status(fiber.StatusCreated).JSON(pkg_success.SuccessCreateData(userResponse))
}

func (h UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	var req user_model.UpdateUserRequest
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id")))
	}

	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Validate
	err = h.xValidator.Validate(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// Call Service
	userResponse, err := h.userService.UpdateUser(ctx.Context(), id, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_success.SuccessCreateData(userResponse))
}

func (h UserHandler) ChangePassword(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id")))
	}

	var req user_model.ChangePasswordRequest
	err = ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	//Validate
	err = h.xValidator.Validate(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	// Call Service
	err = h.userService.ChangePassword(ctx.Context(), id, req)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_success.SuccessGetData(id))
}

func (h UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(pkg_error.NewBadRequest(fmt.Errorf("invalid id")))
	}

	// Call Service
	err = h.userService.DeleteUser(ctx.Context(), id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(fiber.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(fiber.StatusOK).JSON(pkg_success.SuccessDeleteData(id))
}
