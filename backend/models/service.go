package models

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Rate      string    `json:"rate"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Services []*Service
