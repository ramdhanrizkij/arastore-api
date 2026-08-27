package integration

import (
	"context"
	"testing"

	"github.com/ramdhanrizkij/arastore-api/internal/features/category/repository"
	"github.com/ramdhanrizkij/arastore-api/internal/model"
	"github.com/ramdhanrizkij/arastore-api/internal/shared/pagination"
	"github.com/ramdhanrizkij/arastore-api/test/integration/testhelper"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCategoryRepository_FindAll(t *testing.T) {
	tdb := testhelper.SetupTestDB(t)
	defer tdb.Teardown(t)

	repo := repository.NewCategoryRepository(tdb.DB)

	var sampleCategories = []model.Category{
		{Name: "Electronics", Description: "Electronic devices"},
		{Name: "Fashion", Description: "Clothing and accessories"},
		{Name: "Food & Beverage", Description: "Food and drinks"},
		{Name: "Home & Living", Description: "Home and living products"},
		{Name: "Sports", Description: "Sports equipment"},
	}

	// Empty data in categories table
	t.Run("returns empty when no data", func(t *testing.T) {
		tdb.TruncateAll(t)

		pq := &pagination.PaginationQuery{
			Page: 1,
			PerPage: 10,
			Sort: "created_at",
			Order: "desc",
		}

		cats, total, err := repo.FindAll(context.Background(), pq)
		assert.NoError(t, err)
		assert.Equal(t, int64(0), total)
		assert.Empty(t, cats)
	})

	// return all categories without search
	t.Run("returns all categories without search", func(t *testing.T) {
		tdb.TruncateAll(t)
		seedCategories(t, tdb.DB, sampleCategories)

		pq := &pagination.PaginationQuery{
			Page: 1, 
			PerPage: 10,
			Sort: "created_at",
			Order: "desc",
		}

		cats, total, err := repo.FindAll(context.Background(), pq)

		assert.NoError(t, err)
		assert.Equal(t, int64(5),total)
		assert.Len(t, cats, 5)
		
	})
}

func seedCategories(t *testing.T, db *gorm.DB, categories []model.Category) {
	t.Helper()
	for _, c := range categories {
		if err := db.Create(&c).Error; err != nil {
			t.Fatalf("failed to seed category %s: %s", c.Name, err)
		}
	}
}
