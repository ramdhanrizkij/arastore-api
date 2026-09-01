package permission

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/features/permission/domain"
	permHandler "github.com/ramdhanrizkij/arastore-api/internal/features/permission/handler"
)

// Module wires the permission feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *permHandler.PermissionHTTPHandler
}

// NewModule constructs the permission handler from its service and logger.
func NewModule(svc domain.PermissionService, log *zap.Logger) *Module {
	return &Module{handler: permHandler.NewPermissionHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the permission routes to the given router group.
func (m *Module) RegisterRoutes(router fiber.Router, db *gorm.DB, jwtSecret string) {
	permHandler.RegisterRoutes(router, m.handler, db, jwtSecret)
}
