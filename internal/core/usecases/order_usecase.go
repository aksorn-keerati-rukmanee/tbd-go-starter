package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
)

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}
