package handler

import (
	"errors"

	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
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
		h.handleError(c, err)
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
		h.handleError(c, err)
	}

	return response.Success(c, "Product retrieved successfully", product)
}

func (h *ProductHTTPHandler) handleError(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}
	h.log.Error("unexpected error in product handler", zap.Error(err))
	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}
