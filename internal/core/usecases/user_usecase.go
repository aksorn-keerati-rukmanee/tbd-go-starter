package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
)

type UserUseCase interface {
	CreateUser(user entities.User) error
	GetAllUser() ([]entities.User, error)
}
