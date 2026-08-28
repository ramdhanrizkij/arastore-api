package domain

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

type ProductRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]model.Product, int64, error)
	FindByID(ctx context.Context, id string) (*model.Product, error)
}

type ProductService interface {
	GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]ProductResponse, *response.PaginationMeta, error)
	GetByID(ctx context.Context, id string) (*ProductResponse, error)
}
