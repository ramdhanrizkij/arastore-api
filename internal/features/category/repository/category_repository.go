package repository

import (
	"context"
	"errors"

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

// FindByID implements [domain.CategoryRepository].
func (r *categoryRepository) FindByID(ctx context.Context, id string) (*model.Category, error) {
	var category model.Category

	if err := r.db.Where("id=?", id).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, apperrors.ErrNotFound
		}
	}

	return &category, nil
}

// Create implements [domain.CategoryRepository].
func (r *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Create(category).Error; err != nil {
		return apperrors.WrapError(err, "failed to create category")
	}
	return nil
}

// Update implements [domain.CategoryRepository].
func (r *categoryRepository) Update(ctx context.Context, category *model.Category) error {
	if err := r.db.WithContext(ctx).Save(category).Error; err != nil {
		return apperrors.WrapError(err, "failed to update category")
	}
	return nil
}

// Delete implements [domain.CategoryRepository].
func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id=?", id).Delete(&model.Category{})
	if result.Error != nil {
		return apperrors.WrapError(result.Error, "failed to delete category")
	}

	if result.RowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}

// FindByName implements [domain.CategoryRepository].
func (r *categoryRepository) FindByName(ctx context.Context, name string) (*model.Category, error) {
	var category model.Category
	if err := r.db.Where("name=?", name).First(&category).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil,apperrors.ErrNotFound
		}
		return nil,err
	}
	
	return &category, nil
}
