package models

import (
	"github.com/google/uuid"
)

type Service struct {
	ID               uuid.UUID `json:"id,omitempty"`
	Name             string    `json:"name"`
	Price            float32   `json:"price"`
	Rate             string    `json:"rate"`
	RouterIdentifier string    `json:"router_identifier,omitempty"`
}

type Services []*Service
