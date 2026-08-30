package domain

import (
	"time"

	"github.com/google/uuid"

	permissionDomain "github.com/ramdhanrizkij/arastore-api/internal/features/permission/domain"
)

// Role represents a user role (e.g. superadmin, admin, user).
type Role struct {
	ID          uuid.UUID                 `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	Name        string                    `gorm:"type:varchar(50);uniqueIndex;not null" json:"name" example:"manager"`
	Description string                    `gorm:"type:text" json:"description" example:"Can manage operational resources"`
	GuardName   string                    `gorm:"type:varchar(50);default:'api'" json:"guard_name" example:"api"`
	Permissions []permissionDomain.Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
	CreatedAt   time.Time                 `json:"created_at" example:"2026-05-13T10:00:00+07:00"`
	UpdatedAt   time.Time                 `json:"updated_at" example:"2026-05-13T10:00:00+07:00"`
}

// TableName returns the table name for the Role model.
func (Role) TableName() string {
	return "roles"
}

// RolePermission is the explicit join table for the many-to-many relationship
// between roles and permissions.
type RolePermission struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	RoleID       uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_permission" json:"role_id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	PermissionID uuid.UUID `gorm:"type:uuid;not null;uniqueIndex:idx_role_permission" json:"permission_id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	CreatedAt    time.Time `json:"created_at" example:"2026-05-13T10:00:00+07:00"`
}

// TableName returns the table name for the RolePermission model.
func (RolePermission) TableName() string {
	return "role_permissions"
}
