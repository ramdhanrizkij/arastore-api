package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/ramdhanrizkij/arastore-api/internal/features/role/domain"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
)

type roleRepository struct {
	db *gorm.DB
}

// NewRoleRepository creates a RoleRepository backed by GORM.
func NewRoleRepository(db *gorm.DB) domain.RoleRepository {
	return &roleRepository{db: db}
}

func (r *roleRepository) FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]domain.Role, int64, error) {
	var roles []domain.Role
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Role{})

	if pq.Search != "" {
		query = query.Where("name ILIKE ?", "%"+pq.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, apperrors.WrapError(err, "failed to count roles")
	}

	if err := query.
		Order(pq.GetSort()).
		Limit(pq.GetLimit()).
		Offset(pq.GetOffset()).
		Find(&roles).Error; err != nil {
		return nil, 0, apperrors.WrapError(err, "failed to fetch roles")
	}

	return roles, total, nil
}

func (r *roleRepository) FindByID(ctx context.Context, id string) (*domain.Role, error) {
	var role domain.Role
	result := r.db.WithContext(ctx).
		Preload("Permissions").
		Where("id = ?", id).
		First(&role)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, apperrors.WrapError(result.Error, "failed to find role")
	}
	return &role, nil
}

func (r *roleRepository) FindByName(ctx context.Context, name string) (*domain.Role, error) {
	var role domain.Role
	result := r.db.WithContext(ctx).Where("name = ?", name).First(&role)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, apperrors.WrapError(result.Error, "failed to find role by name")
	}
	return &role, nil
}

func (r *roleRepository) Create(ctx context.Context, role *domain.Role) error {
	if err := r.db.WithContext(ctx).Create(role).Error; err != nil {
		return apperrors.WrapError(err, "failed to create role")
	}
	return nil
}

func (r *roleRepository) Update(ctx context.Context, role *domain.Role) error {
	if err := r.db.WithContext(ctx).Save(role).Error; err != nil {
		return apperrors.WrapError(err, "failed to update role")
	}
	return nil
}

func (r *roleRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Role{})
	if result.Error != nil {
		return apperrors.WrapError(result.Error, "failed to delete role")
	}
	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}

// AssignPermissions batch-inserts role_permission rows, skipping duplicates
// via ON CONFLICT DO NOTHING (handled by GORM's clause.OnConflict).
func (r *roleRepository) AssignPermissions(ctx context.Context, roleID string, permissionIDs []string) error {
	roleUUID, err := uuid.Parse(roleID)
	if err != nil {
		return apperrors.BadRequestWith("invalid role ID", err)
	}

	rows := make([]domain.RolePermission, 0, len(permissionIDs))
	for _, pid := range permissionIDs {
		permUUID, err := uuid.Parse(pid)
		if err != nil {
			return apperrors.BadRequestWith("invalid permission ID: "+pid, err)
		}
		rows = append(rows, domain.RolePermission{
			RoleID:       roleUUID,
			PermissionID: permUUID,
		})
	}

	if err := r.db.WithContext(ctx).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&rows).Error; err != nil {
		return apperrors.WrapError(err, "failed to assign permissions")
	}
	return nil
}

// RemovePermissions deletes the specified permission assignments for a role.
func (r *roleRepository) RemovePermissions(ctx context.Context, roleID string, permissionIDs []string) error {
	if err := r.db.WithContext(ctx).
		Where("role_id = ? AND permission_id IN ?", roleID, permissionIDs).
		Delete(&domain.RolePermission{}).Error; err != nil {
		return apperrors.WrapError(err, "failed to remove permissions")
	}
	return nil
}
