package example_http

import (
	"errors"
	"net/http"
	"strconv"

	example_model "github.com/verlinof/fiber-project-structure/internal/modules/example/model"
	pkg_error "github.com/verlinof/fiber-project-structure/pkg/error"
	pkg_success "github.com/verlinof/fiber-project-structure/pkg/success"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h ExampleHandler) GetAll(ctx *fiber.Ctx) error {
	items, err := h.exampleService.GetAll(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(pkg_success.SuccessGetData(items))
}

func (h ExampleHandler) GetByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(errors.New("invalid id")))
	}

	item, err := h.exampleService.GetByID(ctx.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(pkg_success.SuccessGetData(item))
}

func (h ExampleHandler) Create(ctx *fiber.Ctx) error {
	var req example_model.CreateExampleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	item, err := h.exampleService.Create(ctx.Context(), req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(http.StatusCreated).JSON(pkg_success.SuccessCreateData(item))
}

func (h ExampleHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(errors.New("invalid id")))
	}

	var req example_model.UpdateExampleRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	if err := h.xValidator.Validate(req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(err))
	}

	item, err := h.exampleService.Update(ctx.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(pkg_success.SuccessGetData(item))
}

func (h ExampleHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(pkg_error.NewBadRequest(errors.New("invalid id")))
	}

	if err := h.exampleService.Delete(ctx.Context(), id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(http.StatusNotFound).JSON(pkg_error.NewNotFound(err))
		}
		return ctx.Status(http.StatusInternalServerError).JSON(pkg_error.NewInternalServerError(err))
	}

	return ctx.Status(http.StatusOK).JSON(pkg_success.SuccessCreateData("item deleted successfully"))
}
