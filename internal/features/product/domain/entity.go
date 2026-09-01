package domain

import (
	"time"

	"github.com/google/uuid"

	categoryDomain "github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
)

type ProductStatus string

const (
	ProductStatusDraft    ProductStatus = "DRAFT"
	ProductStatusActive   ProductStatus = "ACTIVE"
	ProductStatusInactive ProductStatus = "INACTIVE"
	ProductStatusArchived ProductStatus = "ARCHIVED"
)

type Product struct {
	ID          uuid.UUID                `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id" format:"uuid" example:"018f7606-a3f7-7c40-8e4b-2d47c6e04c8d"`
	CategoryID  uuid.UUID                `gorm:"type:uuid;not nul" json:"category_id"`
	Category    *categoryDomain.Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	SKU         string                   `gorm:"type:varchar(100);uniqueIndex" json:"sku"`
	Name        string                   `gorm:"type:varchar(100)" json:"name"`
	Description string                   `gorm:"type:text" json:"description"`
	Price       float64                  `gorm:"type:numeric(18,2);not null" json:"price"`
	Stock       int                      `gorm:"type:int; not null; default:0" json:"stock"`
	Weight      float64                  `gorm:"type:numeric(10,2)" json:"weight"`
	Status      ProductStatus            `gorm:"type:product_status;not null; default:'DRAFT'" json:"status"`
	CreatedAt   time.Time                `json:"created_at" example:"2026-05-13T10:00:00+07:00"`
	UpdatedAt   time.Time                `json:"updated_at" example:"2026-05-13T10:00:00+07:00"`
}

func (Product) TableName() string {
	return "products"
}
