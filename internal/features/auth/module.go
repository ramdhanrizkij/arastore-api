package auth

import (
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"

	"github.com/ramdhanrizkij/arastore-api/internal/features/auth/domain"
	authHandler "github.com/ramdhanrizkij/arastore-api/internal/features/auth/handler"
)

// Module wires the auth feature's HTTP handler and self-registers its routes.
type Module struct {
	handler *authHandler.AuthHTTPHandler
}

// NewModule constructs the auth handler from its service and logger.
func NewModule(svc domain.AuthService, log *zap.Logger) *Module {
	return &Module{handler: authHandler.NewAuthHTTPHandler(svc, log)}
}

// RegisterRoutes attaches the auth routes to the given router group.
// Auth routes are intentionally unauthenticated.
func (m *Module) RegisterRoutes(router fiber.Router) {
	authHandler.RegisterRoutes(router, m.handler)
}
