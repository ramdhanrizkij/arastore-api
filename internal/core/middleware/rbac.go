package middleware

import (
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"

	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

// permissionCache stores permissions for a role name.
// Key: roleName, Value: cachedPermissions
var (
	permCache sync.Map
	cacheTTL  = 5 * time.Minute
)

type cachedPermissions struct {
	permissions []string
	expiry      time.Time
}

// RequireRole ensures the authenticated user has one of the specified roles.
func RequireRole(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims := GetCurrentUser(c)
		if claims == nil {
			return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
		}

		for _, role := range roles {
			if strings.EqualFold(claims.RoleName, role) {
				return c.Next()
			}
		}

		return response.Error(c, fiber.StatusForbidden, "insufficient role permissions")
	}
}

// RequirePermission ensures the authenticated user's role has one of the specified permissions.
func RequirePermission(db *gorm.DB, permissions ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims := GetCurrentUser(c)
		if claims == nil {
			return response.Error(c, fiber.StatusUnauthorized, "unauthorized")
		}

		// Superadmin bypass
		if strings.EqualFold(claims.RoleName, "superadmin") {
			return c.Next()
		}

		rolePerms, err := getRolePermissions(db, claims.RoleName)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "failed to verify permissions")
		}

		permMap := make(map[string]bool)
		for _, p := range rolePerms {
			permMap[p] = true
		}

		for _, required := range permissions {
			if permMap[required] {
				return c.Next()
			}
		}

		return response.Error(c, fiber.StatusForbidden, "insufficient permissions")
	}
}

func getRolePermissions(db *gorm.DB, roleName string) ([]string, error) {
	// Check cache
	if val, ok := permCache.Load(roleName); ok {
		cp := val.(cachedPermissions)
		if time.Now().Before(cp.expiry) {
			return cp.permissions, nil
		}
	}

	// Cache miss or expired - resolve permission names via a JOIN without
	// depending on any entity struct, keeping this layer free of feature types.
	var permNames []string
	err := db.Table("permissions AS p").
		Joins("JOIN role_permissions AS rp ON rp.permission_id = p.id").
		Joins("JOIN roles AS r ON r.id = rp.role_id").
		Where("r.name = ?", roleName).
		Pluck("p.name", &permNames).Error
	if err != nil {
		return nil, err
	}
	if permNames == nil {
		permNames = []string{}
	}

	// Update cache
	permCache.Store(roleName, cachedPermissions{
		permissions: permNames,
		expiry:       time.Now().Add(cacheTTL),
	})

	return permNames, nil
}

// ClearPermissionCache flushes the in-memory role -> permissions cache so that
// role/permission mutations take effect immediately instead of waiting for the
// TTL to expire. Call it after any role or permission change.
func ClearPermissionCache() {
	permCache.Range(func(k, _ interface{}) bool {
		permCache.Delete(k)
		return true
	})
}
