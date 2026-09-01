package handler

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"github.com/ramdhanrizkij/arastore-api/internal/features/auth/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/validator"
)

// AuthHTTPHandler handles HTTP requests for the auth feature.
// It is the only layer allowed to import Fiber.
type AuthHTTPHandler struct {
	service domain.AuthService
	log     *zap.Logger
}

// NewAuthHTTPHandler creates a new AuthHTTPHandler.
func NewAuthHTTPHandler(service domain.AuthService, log *zap.Logger) *AuthHTTPHandler {
	return &AuthHTTPHandler{service: service, log: log}
}

// Register handles user registration.
func (h *AuthHTTPHandler) Register(c fiber.Ctx) error {
	var req domain.RegisterRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	resp, err := h.service.Register(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Created(c, "User registered successfully", resp)
}

// Login handles user authentication.
func (h *AuthHTTPHandler) Login(c fiber.Ctx) error {
	var req domain.LoginRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	resp, err := h.service.Login(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Login successful", resp)
}

// Refresh handles token refresh.
func (h *AuthHTTPHandler) Refresh(c fiber.Ctx) error {
	var req domain.RefreshTokenRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	resp, err := h.service.Refresh(c.Context(), &req)
	if err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Token refreshed successfully", resp)
}

// Logout handles user logout.
func (h *AuthHTTPHandler) Logout(c fiber.Ctx) error {
	var req domain.LogoutRequest
	if err := c.Bind().Body(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid request body")
	}
	if errs := validator.Validate(req); errs != nil {
		return response.ValidationError(c, errs)
	}

	if err := h.service.Logout(c.Context(), &req); err != nil {
		return h.handleError(c, err)
	}

	return response.Success(c, "Logout successful", nil)
}

// handleError maps AppError values to appropriate HTTP responses.
func (h *AuthHTTPHandler) handleError(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}

	h.log.Error("unexpected error in auth handler", zap.Error(err))
	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}
