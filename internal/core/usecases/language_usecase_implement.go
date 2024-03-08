package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/repositories"
)

type LanguageService struct {
	repo repositories.LanguageRepository
}

func NewLanguageService(repo repositories.LanguageRepository) LanguageUseCase {
	return &LanguageService{repo: repo}
}

func (s *LanguageService) CreateLanguage(Language entities.Language) error {
	return s.repo.Save(Language)
}

func (s *LanguageService) GetAllLanguage() ([]entities.Language, error) {
	// return s.repo.Save(Language)
	return nil, nil
}
