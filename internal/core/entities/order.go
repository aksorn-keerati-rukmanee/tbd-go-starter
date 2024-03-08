package entities

import "github.com/google/uuid"

type Order struct {
	ID    uuid.UUID
	Total float64
}
