package handler

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
)

// RegisterCategoryRoutes registers all categories routes on the provided router group.
func RegisterCategoryRoutes(router fiber.Router, handler *CategoryHTTPHandler, db *gorm.DB, jwtSecret string) {
	categories := router.Group("/categories", middleware.JWTAuth(jwtSecret))

	categories.Get("/", middleware.RequirePermission(db, "categories.view"), handler.GetAll)
	categories.Get("/:id", middleware.RequirePermission(db, "categories.view"), handler.GetByID)
	categories.Post("/",handler.Create)
	categories.Put("/:id",handler.Update)
	categories.Delete("/:id",handler.Delete)
}
