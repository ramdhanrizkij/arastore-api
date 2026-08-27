package domain

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, pq *pagination.PaginationQuery)([]model.Category, int64, error)
}

type CategoryService interface {
}