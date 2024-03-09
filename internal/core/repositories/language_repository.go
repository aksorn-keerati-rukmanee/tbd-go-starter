package repositories

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"

type LanguageRepository interface {
	Create(language entities.Language) error
	All() ([]entities.Language, error)
}
