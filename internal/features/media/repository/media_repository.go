package repository

import (
	"context"

	"github.com/ramdhanrizkij/arastore-api/internal/features/media/domain"
	"gorm.io/gorm"
)

type mediaRepository struct {
	db *gorm.DB
}

func NewMediaRepository(db *gorm.DB) domain.MediaRepository {
	return &mediaRepository{db: db}
}

// Create implements [domain.MediaRepository].
func (r *mediaRepository) Create(ctx context.Context, media *domain.Media) error {
	return r.db.WithContext(ctx).Create(&media).Error
}
