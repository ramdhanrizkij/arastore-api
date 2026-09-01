package domain

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

// PermissionRepository defines the data-access contract for the permission feature.
type PermissionRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]Permission, int64, error)
	FindByID(ctx context.Context, id string) (*Permission, error)
	FindByName(ctx context.Context, name string) (*Permission, error)
	FindByIDs(ctx context.Context, ids []string) ([]Permission, error)
	Create(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
}

// PermissionService defines the business-logic contract for the permission feature.
type PermissionService interface {
	GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]PermissionDetailResponse, *response.PaginationMeta, error)
	GetByID(ctx context.Context, id string) (*PermissionDetailResponse, error)
	Create(ctx context.Context, req *CreatePermissionRequest) (*PermissionDetailResponse, error)
	Update(ctx context.Context, id string, req *UpdatePermissionRequest) (*PermissionDetailResponse, error)
	Delete(ctx context.Context, id string) error
}
