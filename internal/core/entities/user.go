package entities

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Name      string
	Languages []*Language `gorm:"many2many:user_languages;"`
}
