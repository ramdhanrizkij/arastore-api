package seeder

import (
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"gorm.io/gorm"
)

func SeedAddresses(db *gorm.DB) error {
	var user model.User
	if err := db.Where("email=?", "budi@arastore.id").First(&user).Error; err != nil {
		return err		
	}

	addresses := []model.Address{
		{
			UserID:        user.ID,
			Label:         "Rumah",
			RecipientName: "Budi Santoso",
			Phone:         "081234567890",
			Address:       "Jl. Sudirman No. 123, RT 01/RW 05",
			City:          "Jakarta Selatan",
			Province:      "DKI Jakarta",
			PostalCode:    "12190",
		},
		{
			UserID:        user.ID,
			Label:         "Kantor",
			RecipientName: "Budi Santoso",
			Phone:         "081234567890",
			Address:       "Jl. Gatot Subroto No. 45, Gedung Office Tower Lt. 8",
			City:          "Jakarta Pusat",
			Province:      "DKI Jakarta",
			PostalCode:    "10220",
		},
	}

	for _, addr := range addresses {
		var count int64
		db.Model(&model.Address{}).Where("user_id=? AND label=?", addr.UserID, addr.Label).Count(&count)
		if count == 0 {
			if err := db.Create(&addr).Error; err!=nil{
				return err
			}
		}
	}
	return nil
}