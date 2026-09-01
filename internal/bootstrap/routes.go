package bootstrap

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/static"
	"go.uber.org/zap"

	"github.com/ramdhanrizkij/arastore-api/internal/core/config"
	"github.com/ramdhanrizkij/arastore-api/internal/core/storage"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

// SetupRoutes registers the static storage route, the versioned API group, the
// health endpoint, every feature module's routes, and the 404 catch-all.
func SetupRoutes(app *fiber.App, deps *Dependencies) {
	if deps.Storage.ProviderName() == storage.ProviderLocal {
		app.Get(storageBaseURL(deps.Config)+"/*", static.New(storageLocalPath(deps.Config)))
	}

	api := app.Group("/api/v1")
	api.Get("/health", healthCheck(deps))

	deps.AuthModule.RegisterRoutes(api)
	deps.RoleModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)
	deps.PermissionModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)
	deps.UserModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)
	deps.CategoryModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)
	deps.ProductModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)
	deps.MediaModule.RegisterRoutes(api, deps.DB, deps.JWTSecret)

	app.Use(func(c fiber.Ctx) error {
		return response.Error(c, fiber.StatusNotFound, "route not found")
	})
}

func healthCheck(deps *Dependencies) fiber.Handler {
	return func(c fiber.Ctx) error {
		dbStatus := "up"
		cacheStatus := "disabled"

		if deps.Cache.IsEnabled() {
			cacheStatus = "up"
		}

		sqlDB, err := deps.DB.DB()
		if err != nil {
			deps.Logger.Error("failed to get sql.DB for health check", zap.Error(err))
			dbStatus = "down"
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			if err := sqlDB.PingContext(ctx); err != nil {
				deps.Logger.Error("database ping failed during health check", zap.Error(err))
				dbStatus = "down"
			}
		}

		status := "ok"
		message := "service is healthy"
		if dbStatus != "up" {
			status = "degraded"
			message = "service is unhealthy"
		}

		return response.Success(c, message, fiber.Map{
			"status":      status,
			"service":     deps.Config.App.Name,
			"environment": deps.Config.App.Env,
			"database":    dbStatus,
			"cache":       cacheStatus,
			"storage":     deps.Storage.ProviderName(),
		})
	}
}

func storageBaseURL(cfg *config.Config) string {
	if cfg.Storage.BaseURL == "" {
		return "/storage"
	}
	return cfg.Storage.BaseURL
}

func storageLocalPath(cfg *config.Config) string {
	if cfg.Storage.LocalPath == "" {
		return "storage"
	}
	return cfg.Storage.LocalPath
}
