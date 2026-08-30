package user

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/features/user/domain"
	userHandler "github.com/ramdhanrizkij/arastore-api/internal/features/user/handler"
)

// Module wires the user feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *userHandler.UserHTTPHandler
}

// NewModule constructs the user handler from its service and logger.
func NewModule(svc domain.UserService, log *zap.Logger) *Module {
	return &Module{handler: userHandler.NewUserHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the user routes to the given router group.
func (m *Module) RegisterRoutes(router fiber.Router, db *gorm.DB, jwtSecret string) {
	userHandler.RegisterRoutes(router, m.handler, db, jwtSecret)
}
