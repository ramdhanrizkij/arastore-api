package domain

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

type ProductRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]Product, int64, error)
	FindByID(ctx context.Context, id string) (*Product, error)
	FindBySKU(ctx context.Context, sku string) (*Product, error)
	Create(ctx context.Context, product *Product) error
	Update(ctx context.Context, product *Product) error
	Delete(ctx context.Context, id string) error
}

type ProductService interface {
	GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]ProductResponse, *response.PaginationMeta, error)
	GetByID(ctx context.Context, id string) (*ProductResponse, error)
	Create(ctx context.Context, req *CreateProductRequest) (*ProductResponse, error)
	Update(ctx context.Context, id string, req *UpdateProductRequest) (*ProductResponse, error)
	Delete(ctx context.Context, id string) error
}
