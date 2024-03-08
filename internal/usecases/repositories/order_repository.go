package usecases

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
