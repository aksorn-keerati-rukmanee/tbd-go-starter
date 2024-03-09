package repo

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/repositories"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GormUserRepository{db: db}
}

func (r *GormUserRepository) Create(user entities.User) error {
	return r.db.Create(&user).Error
}

func (r *GormUserRepository) All() (result []entities.User, err error) {
	tx := r.db.Preload("Languages").Find(&result)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return result, nil
}
