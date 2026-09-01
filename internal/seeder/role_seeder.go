package seeder

import (
	roleDomain "github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	roles := []roleDomain.Role{
		{Name: "superadmin", Description: "Super Administrator with full access"},
		{Name: "admin", Description: "Administrator"},
		{Name: "user", Description: "Regular user"},
	}

	for _, role := range roles {
		if err := db.Where("name = ?", role.Name).FirstOrCreate(&role).Error; err != nil {
			return err
		}
	}
	return nil
}
