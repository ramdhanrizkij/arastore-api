package repository

import (
	"context"
	"errors"

	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &productRepository{db: db}
}

// FindAll implements [domain.ProductRepository].
func (r *productRepository) FindAll(ctx context.Context, pq *pagination.PaginationQuery) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.WithContext(ctx).Model(&model.Product{})

	if pq.Search != "" {
		query = query.Where("name ILIKE ? OR sku ILIKE ?", "%"+pq.Search+"%", "%"+pq.Search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, apperrors.WrapError(err, "cannot count the product")
	}

	if err := query.Order(pq.GetSort()).
		Limit(pq.GetLimit()).
		Offset(pq.GetOffset()).
		Preload("Category").
		Find(&products).Error; err != nil {
		return nil, 0, apperrors.WrapError(err, "failed to fetch products")
	}

	return products, total, nil
}

// FindByID implements [domain.ProductRepository].
func (r *productRepository) FindByID(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	if err := r.db.Where("id = ?", id).Preload("Category").First(&product).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &product, nil
}
