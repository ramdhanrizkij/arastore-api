package domain

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery)([]model.Category, int64, error)
	FindByID(ctx context.Context, id string) (*model.Category, error)
	FindByName(ctx context.Context, name string) (*model.Category, error)
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
}

type CategoryService interface {
	GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]CategoryResponse, *response.PaginationMeta, error)
	GetByID(ctx context.Context, id string) (*CategoryResponse, error)
	Create(ctx context.Context, req *CreateCategoryRequest) (*CategoryResponse, error)
	Update(ctx context.Context, id string, req *UpdateCategoryRequest) (*CategoryResponse, error)
	Delete(ctx context.Context, id string) error
}