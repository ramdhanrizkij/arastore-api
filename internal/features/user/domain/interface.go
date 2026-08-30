package domain

import (
	"context"

	permissionDomain "github.com/ramdhanrizkij/arastore-api/internal/features/permission/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

// UserRepository defines the data access contract for users.
type UserRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]User, int64, error)
	FindByID(ctx context.Context, id string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id string) error
	GetPermissions(ctx context.Context, userID string) ([]permissionDomain.Permission, error)
}

// UserService defines the business logic contract for users.
type UserService interface {
	GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]UserDetailResponse, *response.PaginationMeta, error)
	GetByID(ctx context.Context, id string) (*UserDetailResponse, error)
	Create(ctx context.Context, req *CreateUserRequest) (*UserDetailResponse, error)
	Update(ctx context.Context, id string, req *UpdateUserRequest) (*UserDetailResponse, error)
	UpdateProfile(ctx context.Context, id string, req *UpdateProfileRequest) (*UserDetailResponse, error)
	Delete(ctx context.Context, currentUserID string, targetID string) error
	GetPermissions(ctx context.Context, userID string) ([]string, error)
}
