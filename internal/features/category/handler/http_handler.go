package handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/validator"
	"go.uber.org/zap"
)

type CategoryHTTPHandler struct {
	service domain.CategoryService
	log     *zap.Logger
}

func NewCategoryHTTPHandler(service domain.CategoryService, log *zap.Logger) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{service: service, log: log}
}

func (h *CategoryHTTPHandler) GetAll(c fiber.Ctx) error {
	pq := pagination.NewPaginationQuery(c)
	categories, meta, err := h.service.GetAll(c.Context(), pq)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.SuccessWithPagination(c, "Categories retrieved successfully", categories, meta)
}

func (h *CategoryHTTPHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	category, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Category retrieved successfully", category)
}

func (h *CategoryHTTPHandler) Create(c fiber.Ctx) error {
	var req domain.CreateCategoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	category, err := h.service.Create(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Created(c, "Category created successfully", category)
}

func (h *CategoryHTTPHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	var req domain.UpdateCategoryRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	category, err := h.service.Update(c.Context(), id, &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Category updated successfully", category)
}

func (h *CategoryHTTPHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	err := h.service.Delete(c.Context(), id)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Category deleted successfully", nil)
}

func (h *CategoryHTTPHandler) handleError(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}
	h.log.Error("unexpected error in category handler", zap.Error(err))
	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}
