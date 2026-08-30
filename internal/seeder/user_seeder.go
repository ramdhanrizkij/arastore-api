package seeder

import (
	roleDomain "github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	userDomain "github.com/ramdhanrizkij/arastore-api/internal/features/user/domain"
	"github.com/ramdhanrizkij/arastore-api/pkg/hash"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	var userRole roleDomain.Role
	if err := db.Where("name = ?", "user").First(&userRole).Error; err != nil {
		return err
	}

	var adminRole roleDomain.Role
	if err := db.Where("name = ?", "admin").First(&adminRole).Error; err != nil {
		return err
	}

	var superAdminRole roleDomain.Role
	if err := db.Where("name = ?", "admin").First(&superAdminRole).Error; err != nil {
		return err
	}

	hashedPassword, err := hash.HashPassword("password")
	if err != nil {
		return err
	}

	users := []userDomain.User{
		{
			Name:     "Budi Santoso",
			Email:    "budi@arastore.id",
			Password: hashedPassword,
			RoleID:   &userRole.ID,
			IsActive: true,
		},
		{
			Name:     "Admin",
			Email:    "admin@arastore.id",
			Password: hashedPassword,
			RoleID:   &adminRole.ID,
			IsActive: true,
		},
		{
			Name:     "Superadmin",
			Email:    "superadmin@arastore.id",
			Password: hashedPassword,
			RoleID:   &superAdminRole.ID,
			IsActive: true,
		},
	}

	for _, user := range users {
		var count int64
		db.Model(&userDomain.User{}).Where("email = ?", user.Email).Count(&count)
		if count == 0 {
			if err := db.Create(&user).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
