package media

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ramdhanrizkij/arastore-api/internal/core/config"
	"github.com/ramdhanrizkij/arastore-api/internal/core/storage"
	mediaHandler "github.com/ramdhanrizkij/arastore-api/internal/features/media/handler"
	"github.com/ramdhanrizkij/arastore-api/internal/features/media/repository"
	"github.com/ramdhanrizkij/arastore-api/internal/features/media/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Module struct {
	handler *mediaHandler.MediaHTTPHandler
}

func NewModule(db *gorm.DB, store storage.Provider, cfg *config.StorageConfig, log *zap.Logger) *Module {
	mediaRepo := repository.NewMediaRepository(db)
	mediaService := service.NewMediaService(mediaRepo, store, cfg.MediaBucket, log)
	mediaHandler := mediaHandler.NewMediaHTTPHandler(mediaService, log)
	return &Module{handler: mediaHandler}
}

func (m *Module) RegisterRoutes(router fiber.Router, jwtSecret string) {
	media := router.Group("/media")

	// Public routes
	media.Get("/:id", m.handler.GetMediaDetail)

	// Protected routes (JWT required)
	media.Post("/", m.handler.UploadMedia)
	media.Delete("/:id", m.handler.DeleteMedia)
}
