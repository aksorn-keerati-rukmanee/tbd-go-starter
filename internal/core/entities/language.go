package entities

import "github.com/google/uuid"

type Language struct {
	ID    uuid.UUID
	Name  string
	Users []*User `gorm:"many2many:user_languages;"`
}
