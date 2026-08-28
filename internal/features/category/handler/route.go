package handler

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
)

// RegisterCategoryRoutes registers all categories routes on the provided router group.
func RegisterCategoryRoutes(router fiber.Router, handler *CategoryHTTPHandler, db *gorm.DB, jwtSecret string) {
	categories := router.Group("/categories")

	categories.Get("/", handler.GetAll)
	categories.Get("/:id", handler.GetByID)
	categories.Post("/", middleware.JWTAuth(jwtSecret), handler.Create)
	categories.Put("/:id", middleware.JWTAuth(jwtSecret), handler.Update)
	categories.Delete("/:id", middleware.JWTAuth(jwtSecret), handler.Delete)
}
