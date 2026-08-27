package seeder

import (
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"gorm.io/gorm"
)

func SeedProducts(db *gorm.DB) error {
	var electronics model.Category
	if err := db.Where("name = ? ", "Electronics").First(&electronics).Error; err != nil {
		return err
	}

	var fashion model.Category
	if err := db.Where("name = ?", "Fashion").First(&fashion).Error; err != nil {
		return err
	}

	var food model.Category
	if err := db.Where("name = ?", "Food & Beverage").First(&food).Error; err != nil {
		return err
	}

	products := []model.Product{
		{
			CategoryID:  electronics.ID,
			SKU:         "ELEC-001",
			Name:        "Laptop ASUS VivoBook 14",
			Description: "Laptop 14 inch, AMD Ryzen 5, 8GB RAM, 512GB SSD",
			Price:       7499000,
			Stock:       25,
			Weight:      1.5,
			Status:      model.ProductStatusActive,
		},
		{
			CategoryID:  electronics.ID,
			SKU:         "ELEC-002",
			Name:        "Samsung Galaxy Buds2",
			Description: "True wireless earbuds with ANC",
			Price:       899000,
			Stock:       50,
			Weight:      0.1,
			Status:      model.ProductStatusActive,
		},
		{
			CategoryID:  fashion.ID,
			SKU:         "FASH-001",
			Name:        "Kemeja Flannel Pria",
			Description: "Kemeja flannel lengan panjang, bahan premium",
			Price:       249000,
			Stock:       100,
			Weight:      0.3,
			Status:      model.ProductStatusActive,
		},
		{
			CategoryID:  fashion.ID,
			SKU:         "FASH-002",
			Name:        "Sneakers Nike Air Max",
			Description: "Sneakers casual pria, warna putih",
			Price:       1299000,
			Stock:       30,
			Weight:      0.8,
			Status:      model.ProductStatusDraft,
		},
		{
			CategoryID:  food.ID,
			SKU:         "FOOD-001",
			Name:        "Kopi Arabica Gayo 250g",
			Description: "Biji kopi arabica single origin Gayo, roasted medium",
			Price:       85000,
			Stock:       200,
			Weight:      0.25,
			Status:      model.ProductStatusActive,
		},
	}

	for _, product := range products {
		var count int64
		db.Model(&model.Product{}).Where("sku = ? ", product.SKU).Count(&count)
		if count == 0 {
			if err := db.Create(&product).Error; err != nil {
				return err
			}
		}
	}
	return nil
}