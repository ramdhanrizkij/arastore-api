package seeder

import (
	"fmt"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Runner struct {
    Name string
    Fn func(db *gorm.DB) error
}

// Run executes all provided seeders in order
func Run(db *gorm.DB, log *zap.Logger, runners ...Runner) error {
	for _, r := range runners {
		log.Info("seeding", zap.String("table", r.Name))
		if err := r.Fn(db); err != nil {
			return fmt.Errorf("seed %s failed: %w", r.Name, err)
		}
		log.Info("seed completed", zap.String("table", r.Name))
	}
	return nil
}