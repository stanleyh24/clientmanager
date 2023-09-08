package models

import (
	"time"

	"github.com/google/uuid"
)

// datos del usuario
type User struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// lista de usuarios
type Users []*User
