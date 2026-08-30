package response

import "github.com/gofiber/fiber/v3"

// Response is the standard envelope for all non-paginated API responses.
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginatedResponse wraps a list payload together with pagination metadata.
type PaginatedResponse struct {
	Success bool             `json:"success"`
	Message string           `json:"message,omitempty"`
	Data    interface{}      `json:"data,omitempty"`
	Meta    *PaginationMeta  `json:"meta,omitempty"`
}

// PaginationMeta contains the information needed for a client to navigate pages.
type PaginationMeta struct {
	CurrentPage int   `json:"current_page" example:"1"`
	PerPage     int   `json:"per_page" example:"10"`
	TotalItems  int64 `json:"total_items" example:"42"`
	TotalPages  int   `json:"total_pages" example:"5"`
}

// Success sends a 200 OK response with the given data payload.
func Success(c fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created sends a 201 Created response with the given data payload.
func Created(c fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessWithPagination sends a 200 OK response that includes pagination metadata.
func SuccessWithPagination(c fiber.Ctx, message string, data interface{}, pagination *PaginationMeta) error {
	return c.Status(fiber.StatusOK).JSON(PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    pagination,
	})
}

// Error sends a response with the given HTTP status code and error message.
func Error(c fiber.Ctx, code int, message string) error {
	return c.Status(code).JSON(Response{
		Success: false,
		Error:   message,
	})
}

// ValidationError sends a 422 Unprocessable Entity response with validation error details.
func ValidationError(c fiber.Ctx, errors interface{}) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		"success": false,
		"message": "validation error",
		"errors":  errors,
	})
}
