package repository

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}


func NewCategoryRepository(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepository{db: db}
}


// FindAll implements [domain.CategoryRepository].
func (r *categoryRepository) FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]model.Category, int64, error) {
	var categories []model.Category
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Category{})

	if pq.Search != "" {
		query = query.Where("name ILIKE ? ", "%"+pq.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, apperrors.WrapError(err, "failed to count categories")
	}

	if err := query.
		Order(pq.GetSort()).
		Limit(pq.GetLimit()).
		Offset(pq.GetOffset()).
		Find(&categories).Error; err != nil {
			return nil, 0, apperrors.WrapError(err, "failed to fetch categories")
	}
	
	return categories, total, nil
}