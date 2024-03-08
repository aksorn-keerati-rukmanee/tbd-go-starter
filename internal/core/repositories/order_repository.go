package repositories

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
