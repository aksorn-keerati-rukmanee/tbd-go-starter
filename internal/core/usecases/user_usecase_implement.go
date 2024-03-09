package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user entities.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetAllUser() ([]entities.User, error) {
	result, err := s.repo.All()
	if err != nil {
		return nil, err
	}
	return result, nil
}
