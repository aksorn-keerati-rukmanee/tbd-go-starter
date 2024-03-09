package repositories

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"

type UserRepository interface {
	Create(User entities.User) error
	All() ([]entities.User, error)
}
