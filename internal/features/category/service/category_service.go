package service

import (
	"context"
	"errors"

	"github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

type categoryService struct {
	repo domain.CategoryRepository
}

func NewCategoryService(repo domain.CategoryRepository) domain.CategoryService {
	return &categoryService{repo: repo}
}

// GetAll implements [domain.CategoryService].
func (s *categoryService) GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]domain.CategoryResponse, *response.PaginationMeta, error) {
	categories, total, err := s.repo.FindAll(ctx, pq)
	if err != nil {
		return nil, nil, err
	}

	resp := make([]domain.CategoryResponse, 0, len(categories))
	for _, category := range categories {
		resp = append(resp, toCategoryResponse(category))
	}

	meta := &response.PaginationMeta{
		CurrentPage: pq.Page,
		PerPage:     pq.PerPage,
		TotalItems:  total,
		TotalPages:  pagination.CalculateTotalPages(total, pq.PerPage),
	}

	return resp, meta, nil
}

// GetByID implements [domain.CategoryService].
func (s *categoryService) GetByID(ctx context.Context, id string) (*domain.CategoryResponse, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := toCategoryResponse(*category)
	return &resp, nil
}

func toCategoryResponse(r model.Category) domain.CategoryResponse {
	return domain.CategoryResponse{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt.Format("2026-01-02T15:04:05Z07:00"),
		UpdatedAt:   r.UpdatedAt.Format("2026-01-02T15:04:05Z07:00"),
	}
}

// Create implements [domain.CategoryService].
func (s *categoryService) Create(ctx context.Context, req *domain.CreateCategoryRequest) (*domain.CategoryResponse, error) {
	// Check unique name
	existing, err := s.repo.FindByName(ctx, req.Name)
	if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
		return nil, apperrors.WrapError(err, "failed to check category name")
	}

	if existing != nil {
		return nil, apperrors.NewAppError(409, "category name already exists", nil)
	}

	category := &model.Category{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := s.repo.Create(ctx, category); err != nil {
		return nil, err
	}

	resp := toCategoryResponse(*category)
	return &resp, nil

}

// Update implements [domain.CategoryService].
func (s *categoryService) Update(ctx context.Context, id string, req *domain.UpdateCategoryRequest) (*domain.CategoryResponse, error) {
	category, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if category.Name != req.Name {
		existing, err := s.repo.FindByName(ctx, req.Name)
		if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
			return nil, apperrors.WrapError(err, "failed to check category name")
		}

		if existing != nil {
			return nil, apperrors.NewAppError(409, "category name already exists", nil)
		}
	}

	category.Name = req.Name
	category.Description = req.Description

	if err := s.repo.Update(ctx, category); err != nil {
		return nil, err
	}

	resp := toCategoryResponse(*category)
	return &resp, nil
}

// Delete implements [domain.CategoryService].
func (s *categoryService) Delete(ctx context.Context, id string) error {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		return err
	}
	
	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}
	
	return nil
}