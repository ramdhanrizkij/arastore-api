package bootstrap

import (
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/core/cache"
	"github.com/ramdhanrizkij/arastore-api/internal/core/config"
	"github.com/ramdhanrizkij/arastore-api/internal/core/storage"
	"github.com/ramdhanrizkij/arastore-api/internal/core/worker"

	"github.com/ramdhanrizkij/arastore-api/internal/features/auth"
	authRepo "github.com/ramdhanrizkij/arastore-api/internal/features/auth/repository"
	authService "github.com/ramdhanrizkij/arastore-api/internal/features/auth/service"

	"github.com/ramdhanrizkij/arastore-api/internal/features/category"
	categoryRepo "github.com/ramdhanrizkij/arastore-api/internal/features/category/repository"
	categoryService "github.com/ramdhanrizkij/arastore-api/internal/features/category/service"

	"github.com/ramdhanrizkij/arastore-api/internal/features/permission"
	permRepo "github.com/ramdhanrizkij/arastore-api/internal/features/permission/repository"
	permService "github.com/ramdhanrizkij/arastore-api/internal/features/permission/service"

	"github.com/ramdhanrizkij/arastore-api/internal/features/product"
	productRepo "github.com/ramdhanrizkij/arastore-api/internal/features/product/repository"
	productService "github.com/ramdhanrizkij/arastore-api/internal/features/product/service"

	"github.com/ramdhanrizkij/arastore-api/internal/features/role"
	roleRepo "github.com/ramdhanrizkij/arastore-api/internal/features/role/repository"
	roleService "github.com/ramdhanrizkij/arastore-api/internal/features/role/service"

	"github.com/ramdhanrizkij/arastore-api/internal/features/user"
	userRepo "github.com/ramdhanrizkij/arastore-api/internal/features/user/repository"
	userService "github.com/ramdhanrizkij/arastore-api/internal/features/user/service"
)

// Dependencies is the composition root: it holds all wired feature modules plus
// the shared infrastructure needed for route registration and health checks.
type Dependencies struct {
	DB        *gorm.DB
	Cache     cache.Client
	Storage   storage.Provider
	Worker    *worker.WorkerPool
	Logger    *zap.Logger
	Config    *config.Config
	JWTSecret string

	AuthModule       *auth.Module
	RoleModule       *role.Module
	PermissionModule *permission.Module
	UserModule       *user.Module
	CategoryModule   *category.Module
	ProductModule    *product.Module
}

// NewDependencies wires every feature (repository -> service -> module) in a
// single, auditable place.
func NewDependencies(
	cfg *config.Config,
	db *gorm.DB,
	cacheClient cache.Client,
	storageProvider storage.Provider,
	wp *worker.WorkerPool,
	logger *zap.Logger,
) *Dependencies {
	jwtSecret := cfg.JWT.Secret
	cacheTTL := computeCacheTTL(cfg)

	// Auth
	authRepository := authRepo.NewAuthRepository(db)
	authSvc := authService.NewAuthService(authRepository, wp, jwtSecret, cfg.JWT.ExpiryHours, cfg.JWT.RefreshExpiryHours, logger)
	authMod := auth.NewModule(authSvc, logger)

	// Role
	roleRepository := roleRepo.NewRoleRepository(db)
	roleSvc := roleService.NewRoleService(roleRepository, cacheClient, cacheTTL, logger)
	roleMod := role.NewModule(roleSvc, logger)

	// Permission
	permRepository := permRepo.NewPermissionRepository(db)
	permSvc := permService.NewPermissionService(permRepository, cacheClient, cacheTTL, logger)
	permMod := permission.NewModule(permSvc, logger)

	// User
	userRepository := userRepo.NewUserRepository(db)
	userSvc := userService.NewUserService(userRepository, cacheClient, storageProvider, cfg.Storage.DefaultBucket, cacheTTL, logger)
	userMod := user.NewModule(userSvc, logger)

	// Category
	categoryRepository := categoryRepo.NewCategoryRepository(db)
	categorySvc := categoryService.NewCategoryService(categoryRepository)
	categoryMod := category.NewModule(categorySvc, logger)

	// Product
	productRepository := productRepo.NewProductRepository(db)
	productSvc := productService.NewProductService(productRepository)
	productMod := product.NewModule(productSvc, logger)

	return &Dependencies{
		DB:        db,
		Cache:     cacheClient,
		Storage:   storageProvider,
		Worker:    wp,
		Logger:    logger,
		Config:    cfg,
		JWTSecret: jwtSecret,

		AuthModule:       authMod,
		RoleModule:       roleMod,
		PermissionModule: permMod,
		UserModule:       userMod,
		CategoryModule:   categoryMod,
		ProductModule:    productMod,
	}
}

func computeCacheTTL(cfg *config.Config) time.Duration {
	ttlMinutes := cfg.Redis.CacheTTLMinutes
	if ttlMinutes <= 0 {
		ttlMinutes = 5
	}
	return time.Duration(ttlMinutes) * time.Minute
}
