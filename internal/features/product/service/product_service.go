package service

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/features/product/domain"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
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
