package domain

type CreateProductRequest struct {
	CategoryID  string  `json:"category_id" validate:"required,uuid"`
	SKU         string  `json:"sku" validate:"required,min:1,max:100"`
	Name        string  `json:"name" validate:"required,min:1,max:256"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"min:0"`
	Weight      float64 `json:"weight" validate:"min:0"`
	Status      string  `json:"status" validate:"omitempty,oneof=DRAFT ACTIVE INACTIVE ARCHIVED"`
}

type UpdateProductRequest struct {
	CategoryID  string  `json:"category_id" validate:"required,uuid"`
	SKU         string  `json:"sku" validate:"required,min:1,max:100"`
	Name        string  `json:"name" validate:"required,min:1,max:100"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Stock       int     `json:"stock" validate:"min:0"`
	Weight      float64 `json:"weight" validate:"min:0"`
	Status      string  `json:"status" validate:"required,oneof=DRAFT ACTIVE INACTIVE ARCHIVED"`
}

type ProductCategoryResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type ProductResponse struct {
	ID          string                   `json:"id"`
	CategoryID  string                   `json:"category_id"`
	Category    *ProductCategoryResponse `json:"category"`
	SKU         string                   `json:"sku"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Price       float64                  `json:"price"`
	Stock       int                      `json:"stock"`
	Weight      float64                  `json:"weight"`
	Status      string                   `json:"status"`
	CreatedAt   string                   `json:"created_at"`
	UpdatedAt   string                   `json:"updated_at"`
}
