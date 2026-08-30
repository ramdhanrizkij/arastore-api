package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
	"gorm.io/gorm"
)

func RegisterProductRoutes(router fiber.Router, handler *ProductHTTPHandler, db *gorm.DB, jwtSecret string) {
	products := router.Group("/products")

	products.Get("/", handler.GetAll)
	products.Get("/:id", handler.GetByID)
	products.Post("/", handler.Create)
	products.Put("/:id", middleware.JWTAuth(jwtSecret), middleware.RequirePermission(db, "products.edit"), handler.Update)
	products.Delete("/:id", middleware.JWTAuth(jwtSecret), middleware.RequirePermission(db, "products.delete"), handler.Delete)

}
