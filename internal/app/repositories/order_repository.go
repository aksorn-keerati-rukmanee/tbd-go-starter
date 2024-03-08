package repositories

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/app/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
