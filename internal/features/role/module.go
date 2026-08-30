package role

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	roleHandler "github.com/ramdhanrizkij/arastore-api/internal/features/role/handler"
)

// Module wires the role feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *roleHandler.RoleHTTPHandler
}

// NewModule constructs the role handler from its service and logger.
func NewModule(svc domain.RoleService, log *zap.Logger) *Module {
	return &Module{handler: roleHandler.NewRoleHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the role routes (with JWT + permission middleware) to
// the given router group.
func (m *Module) RegisterRoutes(router fiber.Router, db *gorm.DB, jwtSecret string) {
	roleHandler.RegisterRoutes(router, m.handler, db, jwtSecret)
}
