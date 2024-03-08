package adapters

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/repositories"
	"gorm.io/gorm"
)

type GormOrderRepository struct {
	db *gorm.DB
}

func NewGormOrderRepository(db *gorm.DB) repositories.OrderRepository {
	return &GormOrderRepository{db: db}
}

func (r *GormOrderRepository) Save(order entities.Order) error {
	return r.db.Create(&order).Error
}
