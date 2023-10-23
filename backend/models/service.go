package models

import (
	"github.com/google/uuid"
)

type Service struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	Rate      string    `json:"rate"`
	CreatedAt int64     `json:"created_at"`
	UpdatedAt int64     `json:"updated_at"`
}

type Services []*Service
