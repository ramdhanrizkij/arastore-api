package handler

import (
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func RegisterProductRoutes(router fiber.Router, handler *ProductHTTPHandler, db *gorm.DB, jwtSecret string) {
	products := router.Group("/products")

	products.Get("/", handler.GetAll)
	products.Get("/:id", handler.GetByID)
}
