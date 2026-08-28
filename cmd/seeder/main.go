package main

import (
	"log"

	"github.com/ramdhanrizkij/arastore-api/internal/core/config"
	"github.com/ramdhanrizkij/arastore-api/internal/core/database"
	"github.com/ramdhanrizkij/arastore-api/internal/seeder"
	"github.com/ramdhanrizkij/arastore-api/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	if err := logger.InitGlobal(cfg.Log.Level); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	defer logger.Log.Sync()

	db, err := database.NewPostgresDB(&cfg.DB, cfg.App.Env, logger.Log)
	if err != nil {
		logger.Log.Fatal("failed to conect to database", zap.Error(err))
	}

	defer database.CloseDB(db)

	err = seeder.Run(db, logger.Log,
		seeder.Runner{Name: "roles", Fn: seeder.SeedRoles},
		seeder.Runner{Name: "users", Fn: seeder.SeedUsers},
		seeder.Runner{Name: "permissions", Fn: seeder.SeedPermissions},
		seeder.Runner{Name: "addresses", Fn: seeder.SeedAddresses},
		seeder.Runner{Name: "caategories", Fn: seeder.SeedCategories},
		seeder.Runner{Name: "products", Fn: seeder.SeedProducts},
	)

	if err != nil {
		logger.Log.Fatal("seeding failed", zap.Error(err))
	}

	logger.Log.Info("all seeders completed successfully")
}
