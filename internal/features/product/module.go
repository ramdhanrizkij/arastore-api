package product

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	productHandler "github.com/ramdhanrizkij/arastore-api/internal/features/product/handler"
)

// Module wires the product feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *productHandler.ProductHTTPHandler
}

// NewModule constructs the product handler from its service and logger.
func NewModule(svc domain.ProductService, log *zap.Logger) *Module {
	return &Module{handler: productHandler.NewProductHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the product routes to the given router group.
func (m *Module) RegisterRoutes(router fiber.Router, db *gorm.DB, jwtSecret string) {
	productHandler.RegisterProductRoutes(router, m.handler, db, jwtSecret)
}
