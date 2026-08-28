package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	apperrors "github.com/ramdhanrizkij/arastore-api/internal/shared/errors"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/response"
)

type productService struct {
	repo domain.ProductRepository
}

func NewProductService(repo domain.ProductRepository) domain.ProductService {
	return &productService{repo: repo}
}

// GetAll implements [domain.ProductService].
func (s *productService) GetAll(ctx context.Context, pq *pagination.PaginationQuery) ([]domain.ProductResponse, *response.PaginationMeta, error) {
	products, total, err := s.repo.FindAll(ctx, pq)
	if err != nil {
		return nil, nil, err
	}

	resp := make([]domain.ProductResponse, 0, len(products))
	for _, product := range products {
		resp = append(resp, toProductResponse(product))
	}

	meta := &response.PaginationMeta{
		CurrentPage: pq.Page,
		PerPage:     pq.PerPage,
		TotalItems:  total,
		TotalPages:  pagination.CalculateTotalPages(total, pq.PerPage),
	}

	return resp, meta, nil
}

// GetByID implements [domain.ProductService].
func (s *productService) GetByID(ctx context.Context, id string) (*domain.ProductResponse, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	resp := toProductResponse(*product)
	return &resp, err
}

// Create implements [domain.ProductService].
func (s *productService) Create(ctx context.Context, req *domain.CreateProductRequest) (*domain.ProductResponse, error) {
	existing, err := s.repo.FindBySKU(ctx, req.SKU)
	if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
		return nil, apperrors.WrapError(err, "failed to check product sku")
	}

	if existing != nil {
		return nil, apperrors.NewAppError(409, "product with that SKU already exists", nil)
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, apperrors.NewAppError(422, "category_id format not valid", nil)
	}

	status := model.ProductStatus(req.Status)
	if status == "" {
		status = model.ProductStatusDraft
	}

	product := &model.Product{
		Name:        req.Name,
		CategoryID:  categoryID,
		SKU:         req.SKU,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		Weight:      req.Weight,
		Status:      status,
	}

	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	resp := toProductResponse(*product)
	return &resp, nil
}

// Update implements [domain.ProductService].
func (s *productService) Update(ctx context.Context, id string, req *domain.UpdateProductRequest) (*domain.ProductResponse, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product.SKU != req.SKU {
		existing, err := s.repo.FindBySKU(ctx, req.SKU)
		if err != nil && !errors.Is(err, apperrors.ErrNotFound) {
			return nil, apperrors.WrapError(err, "failed to check product sku")
		}

		if existing != nil {
			return nil, apperrors.NewAppError(409, "product sku already exists", nil)
		}
	}

	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return nil, apperrors.NewAppError(422, "category_id format are not valid", nil)
	}

	product.CategoryID = categoryID
	product.SKU = req.SKU
	product.Name = req.Name
	product.Description = req.Description
	product.Price = req.Price
	product.Stock = req.Stock
	product.Weight = req.Weight
	product.Status = model.ProductStatus(req.Status)

	if err := s.repo.Update(ctx, product); err != nil {
		return nil, err
	}

	resp := toProductResponse(*product)
	return &resp, nil
}

// Delete implements [domain.ProductService].
func (s *productService) Delete(ctx context.Context, id string) error {
	if _, err := s.repo.FindByID(ctx, id); err != nil {
		return err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}

func toProductResponse(r model.Product) domain.ProductResponse {
	return domain.ProductResponse{
		ID:         r.ID.String(),
		CategoryID: r.CategoryID.String(),
		Category: &domain.ProductCategoryResponse{
			ID:          r.Category.ID.String(),
			Name:        r.Category.Name,
			Description: r.Category.Description,
		},
		SKU:         r.SKU,
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Stock:       r.Stock,
		Weight:      r.Weight,
		Status:      string(r.Status),
		CreatedAt:   r.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:   r.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
	}
}
