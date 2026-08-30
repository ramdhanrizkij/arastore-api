package handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
	"github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/validator"
)

// RoleHTTPHandler handles HTTP requests for the role feature.
type RoleHTTPHandler struct {
	service domain.RoleService
	log     *zap.Logger
}

// NewRoleHTTPHandler creates a new RoleHTTPHandler.
func NewRoleHTTPHandler(service domain.RoleService, log *zap.Logger) *RoleHTTPHandler {
	return &RoleHTTPHandler{service: service, log: log}
}

// GetAll returns paginated roles.
func (h *RoleHTTPHandler) GetAll(c fiber.Ctx) error {
	pq := pagination.NewPaginationQuery(c)
	roles, meta, err := h.service.GetAll(c.Context(), pq)
	if err != nil {
		return h.handleError(c, err)
	}
	return response.SuccessWithPagination(c, "Roles retrieved successfully", roles, meta)
}

// GetByID returns a role by UUID.
func (h *RoleHTTPHandler) GetByID(c fiber.Ctx) error {
	id := c.Params("id")
	role, err := h.service.GetByID(c.Context(), id)
	if err != nil {
		return h.handleError(c, err)
	}
	return response.Success(c, "Role retrieved successfully", role)
}

// Create creates a new role.
func (h *RoleHTTPHandler) Create(c fiber.Ctx) error {
	var req domain.CreateRoleRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}
	role, err := h.service.Create(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}
	return response.Created(c, "Role created successfully", role)
}

// Update updates an existing role.
func (h *RoleHTTPHandler) Update(c fiber.Ctx) error {
	id := c.Params("id")
	var req domain.UpdateRoleRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}
	role, err := h.service.Update(c.Context(), id, &req)
	if err != nil {
		return h.handleError(c, err)
	}
	return response.Success(c, "Role updated successfully", role)
}

// Delete deletes a role.
func (h *RoleHTTPHandler) Delete(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return h.handleError(c, err)
	}
	return response.Success(c, "Role deleted successfully", nil)
}

// AssignPermissions assigns permissions to a role.
func (h *RoleHTTPHandler) AssignPermissions(c fiber.Ctx) error {
	id := c.Params("id")
	var req domain.AssignPermissionsRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}
	if err := h.service.AssignPermissions(c.Context(), id, &req); err != nil {
		return h.handleError(c, err)
	}
	return response.Success(c, "Permissions assigned to role successfully", nil)
}

// RemovePermissions removes permissions from a role.
func (h *RoleHTTPHandler) RemovePermissions(c fiber.Ctx) error {
	id := c.Params("id")
	var req domain.AssignPermissionsRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}
	if err := h.service.RemovePermissions(c.Context(), id, &req); err != nil {
		return h.handleError(c, err)
	}
	return response.Success(c, "Permissions removed from role successfully", nil)
}

// handleError maps AppError to HTTP response codes.
func (h *RoleHTTPHandler) handleError(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}
	h.log.Error("unexpected error in role handler", zap.Error(err))
	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}

// ensure middleware import is used (GetCurrentUser available for future RBAC use)
var _ = middleware.GetCurrentUser
