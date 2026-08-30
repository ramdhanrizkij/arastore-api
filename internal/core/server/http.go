package server

import (
	"errors"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/bootstrap"
	"github.com/ramdhanrizkij/arastore-api/internal/core/cache"
	"github.com/ramdhanrizkij/arastore-api/internal/core/config"
	"github.com/ramdhanrizkij/arastore-api/internal/core/middleware"
	"github.com/ramdhanrizkij/arastore-api/internal/core/storage"
	"github.com/ramdhanrizkij/arastore-api/internal/core/worker"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

// Server represents the HTTP server container. It owns the Fiber app, the wired
// feature dependencies, and the server lifecycle. All feature wiring and route
// registration lives in internal/bootstrap.
type Server struct {
	app    *fiber.App
	deps   *bootstrap.Dependencies
	config *config.Config
	logger *zap.Logger
}

// NewServer initializes the Fiber application, registers global middleware,
// builds shared infrastructure (cache, storage) and the dependency graph, and
// returns a Server ready to have its routes set up.
func NewServer(cfg *config.Config, db *gorm.DB, logger *zap.Logger, wp *worker.WorkerPool, sched *worker.Scheduler) (*Server, error) {
	app := fiber.New(fiber.Config{
		AppName:      cfg.App.Name,
		BodyLimit:    4 * 1024 * 1024, // 4MB
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorHandler: customErrorHandler,
	})

	// Global Middleware
	app.Use(recover.New())
	app.Use(cors.New())
	app.Use(middleware.RequestLogger(logger))

	cacheClient, err := cache.NewClient(&cfg.Redis, logger)
	if err != nil {
		return nil, err
	}

	storageProvider, err := storage.NewProvider(&cfg.Storage, logger)
	if err != nil {
		return nil, err
	}

	deps := bootstrap.NewDependencies(cfg, db, cacheClient, storageProvider, wp, logger)

	return &Server{
		app:    app,
		deps:   deps,
		config: cfg,
		logger: logger,
	}, nil
}

// SetupRoutes registers all application routes via the bootstrap composition root.
func (s *Server) SetupRoutes() {
	bootstrap.SetupRoutes(s.app, s.deps)
}

// Start runs the HTTP server.
func (s *Server) Start() error {
	addr := fmt.Sprintf(":%d", s.config.App.Port)
	s.logger.Info("server starting", zap.Int("port", s.config.App.Port))
	return s.app.Listen(addr)
}

// Shutdown gracefully stops the HTTP server and closes shared infrastructure.
func (s *Server) Shutdown() error {
	s.logger.Info("shutting down HTTP server...")
	if err := s.app.Shutdown(); err != nil {
		return err
	}
	if err := s.deps.Storage.Close(); err != nil {
		return err
	}
	return s.deps.Cache.Close()
}

// AppForTest returns the underlying Fiber app instance for testing purposes.
func (s *Server) AppForTest() *fiber.App {
	return s.app
}

// customErrorHandler converts AppError into the standard API response format.
func customErrorHandler(c fiber.Ctx, err error) error {
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		return response.Error(c, appErr.Code, appErr.Message)
	}

	return response.Error(c, fiber.StatusInternalServerError, "internal server error")
}
