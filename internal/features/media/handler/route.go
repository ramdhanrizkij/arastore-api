package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
	"gorm.io/gorm"
)

func RegisterMediaRoutes(router fiber.Router, handler *MediaHTTPHandler, db *gorm.DB, jwtSecret string) {
	media := router.Group("/media")

	media.Post("/", middleware.JWTAuth(jwtSecret), handler.UploadMedia)
	media.Delete("/:id", middleware.JWTAuth(jwtSecret), handler.DeleteMedia)
	media.Get("/:id", handler.GetMediaDetail)
}
