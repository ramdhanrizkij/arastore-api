package seeder

import (
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"gorm.io/gorm"
)

func SeedPermissions(db *gorm.DB) error {
	permissions := []model.Permission{
		{Name: "roles.view", Description: "View roles list and details"},
		{Name: "roles.create", Description: "Create new roles"},
		{Name: "roles.edit", Description: "Update existing roles"},
		{Name: "roles.delete", Description: "Delete roles"},
		{Name: "roles.assign-permission", Description: "Assign permissions to roles"},
		{Name: "roles.remove-permission", Description: "Remove permissions from roles"},
		{Name: "permissions.view", Description: "View permissions list and details"},
		{Name: "permissions.create", Description: "Create new permissions"},
		{Name: "permissions.edit", Description: "Update existing permissions"},
		{Name: "permissions.delete", Description: "Delete permissions"},
		{Name: "users.view", Description: "View users list and details"},
		{Name: "users.create", Description: "Create new users"},
		{Name: "users.edit", Description: "Update existing users"},
		{Name: "users.delete", Description: "Delete users"},
		{Name: "categories.view", Description: "View category list and details"},
		{Name: "categories.create", Description: "Create category"},
		{Name: "categories.edit", Description: "Edit Category"},
		{Name: "categories.delete", Description: "Delete Category"},
		{Name: "products.view", Description: "View product list and details"},
		{Name: "products.create", Description: "Create Product"},
		{Name: "products.edit", Description: "Edit Product"},
		{Name: "products.delete", Description: "Delete Product"},
	}

	for _, perm := range permissions {
		if err := db.Where("name = ?", perm.Name).FirstOrCreate(&perm).Error; err != nil {
			return err
		}
	}

	// Assign all permissions to superadmin
	var superadmin model.Role
	if err := db.Where("name = ?", "superadmin").First(&superadmin).Error; err != nil {
		return err
	}

	var allPermissions []model.Permission
	if err := db.Find(&allPermissions).Error; err != nil {
		return err
	}

	if err := db.Model(&superadmin).Association("Permissions").Replace(allPermissions); err != nil {
		return err
	}

	// Assign admin permission
	var admin model.Role
	if err := db.Where("name = ?", "admin").First(&admin).Error; err != nil {
		return err
	}

	var adminPermissions []model.Permission
	if err := db.Where("name IN ?", []string{
		"roles.view",
		"permissions.view",
		"users.view",
		"users.create",
		"users.edit",
		"users.delete",
		"categories.view",
		"categories.create",
		"categories.edit",
		"categories.delete",
		"products.view",
		"products.create",
		"products.edit",
		"products.delete",
	}).Find(&adminPermissions).Error; err != nil {
		return err
	}

	// Assign admin permission to role
	if err := db.Model(&admin).Association("Permissions").Replace(adminPermissions); err != nil {
		return err
	}

	// Assign user permission
	var user model.Role
	if err := db.Where("name=?", "user").First(&user).Error; err != nil {
		return err
	}

	var userPermissions []model.Permission
	if err := db.Where("name IN ?", []string{
		"roles.view",
		"permissions.view",
		"users.view",
		"categories.view",
		"products.view",
	}).Find(&userPermissions).Error; err != nil {
		return err
	}

	if err := db.Model(&user).Association("Permissions").Replace(userPermissions); err != nil {
		return err
	}

	return nil
}
