package repo

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/repositories"
	"gorm.io/gorm"
)

type GormLanguageRepository struct {
	db *gorm.DB
}

func NewGormLanguageRepository(db *gorm.DB) repositories.LanguageRepository {
	return &GormLanguageRepository{db: db}
}

func (r *GormLanguageRepository) Save(language entities.Language) error {
	return r.db.Create(&language).Error
}

func (r *GormLanguageRepository) Find() ([]entities.Language, error) {
	// return r.db.Create(&order).Error
	return nil, nil
}
