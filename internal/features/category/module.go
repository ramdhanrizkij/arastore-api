package category

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
	categoryHandler "github.com/ramdhanrizkij/arastore-api/internal/features/category/handler"
)

// Module wires the category feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *categoryHandler.CategoryHTTPHandler
}

// NewModule constructs the category handler from its service and logger.
func NewModule(svc domain.CategoryService, log *zap.Logger) *Module {
	return &Module{handler: categoryHandler.NewCategoryHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the category routes to the given router group.
func (m *Module) RegisterRoutes(router fiber.Router, db *gorm.DB, jwtSecret string) {
	categoryHandler.RegisterCategoryRoutes(router, m.handler, db, jwtSecret)
}
