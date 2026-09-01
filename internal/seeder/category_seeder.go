package seeder

import (
	categoryDomain "github.com/ramdhanrizkij/arastore-api/internal/features/category/domain"
	"gorm.io/gorm"
)

func SeedCategories(db *gorm.DB) error {
	categories := []categoryDomain.Category{
		{Name: "Electronics", Description: "Perangkat elektronik dan gadget"},
		{Name: "Fashion", Description: "Pakaian dan aksesoris fashion"},
		{Name: "Food & Beverage", Description: "Makanan dan minuman"},
		{Name: "Home & Living", Description: "Perlengkapan rumah tangga"},
		{Name: "Sports", Description: "Peralatan olahraga dan outdoor"},
	}

	for _, category := range categories {
		if err := db.Where("name=?", category.Name).FirstOrCreate(&category).Error; err != nil {
			return err
		}
	}
	return nil
}
