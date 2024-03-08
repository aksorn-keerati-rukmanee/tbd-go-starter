package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/repositories"
)

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	repo repositories.OrderRepository
}

func NewOrderService(repo repositories.OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {
	return s.repo.Save(order)
}
