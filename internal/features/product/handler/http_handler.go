package handler

import (
	"errors"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/validator"
	"go.uber.org/zap"
)

type ProductHTTPHandler struct {
	service domain.ProductService
	log     *zap.Logger
}

func NewProductHTTPHandler(service domain.ProductService, log *zap.Logger) *ProductHTTPHandler {
	return &ProductHTTPHandler{
		service: service,
		log:     log,
	}
}

func (h *ProductHTTPHandler) GetAll(c fiber.Ctx) error {
	pq := pagination.NewPaginationQuery(c)
	products, meta, err := h.service.GetAll(c.Context(), pq)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.SuccessWithPagination(c, "Product retrieved successfully", products, meta)
}

func (h *ProductHTTPHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, "product ID is required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid product ID format")
	}
	product, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Product retrieved successfully", product)
}

func (h *ProductHTTPHandler) Create(c fiber.Ctx) error {
	var req domain.CreateProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	product, err := h.service.Create(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Created(c, "Product created successfully", product)
}

func (h *ProductHTTPHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, "product ID are required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid product ID format")
	}

	var req domain.UpdateProductRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	product, err := h.service.Update(c.Context(), id, &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Product updated successfully", product)
}

func (h *ProductHTTPHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, "Product ID are required")
	}

	if _, err := uuid.Parse(id); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid product ID format")
	}

	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Product delete successfully", nil)
}

func (h *ProductHTTPHandler) handleError(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}
	h.log.Error("unexpected error in product handler", zap.Error(err))
	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}
