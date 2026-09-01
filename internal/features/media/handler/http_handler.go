package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
	"go.uber.org/zap"
)

type MediaHTTPHandler struct {
	log *zap.Logger
}

func NewMediaHTTPHandler(log *zap.Logger) *MediaHTTPHandler {
	return &MediaHTTPHandler{log: log}
}

// POST /api/v1/media
// RESPONSE :
//
//	{
//	  "id": "uuid",
//	  "objectKey": "products/2026/09/abc.jpg",
//	  "url": "https://cdn.example.com/abc.jpg"
//	}
func (h *MediaHTTPHandler) UploadMedia(c fiber.Ctx) error {
	// TODO: Implement media upload api
	file, err := c.FormFile("file")
	if err != nil {
		return errors.
	}
	return response.Success(c, "Media uploaded", nil)
}

// GET /api/v1/media/{id}
func (h *MediaHTTPHandler) GetMediaDetail(c fiber.Ctx) error {
	return response.Success(c, "Success get media detail", nil)
}

// DELETE /api/v1/media/{id}
func (h *MediaHTTPHandler) DeleteMedia(c fiber.Ctx) error {
	return response.Success(c, "Success delete media", nil)
}
