package repositories

import "github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"

type LanguageRepository interface {
	Save(language entities.Language) error
	Find() ([]entities.Language, error)
}
