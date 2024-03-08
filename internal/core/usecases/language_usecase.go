package usecases

import (
	"github.com/aksorn-keerati-rukmanee/tbd-go-starter/internal/core/entities"
)

type LanguageUseCase interface {
	CreateLanguage(language entities.Language) error
	GetAllLanguage() ([]entities.Language, error)
}
