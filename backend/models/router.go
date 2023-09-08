package models

import (
	"github.com/google/uuid"
)

type Router struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Ip        string    `json:"ip"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt int64     ` json:"created_at"`
	UpdatedAt int64     ` json:"updated_at"`
}

type Routers []*Router
