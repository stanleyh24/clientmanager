package models

import (
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name"`
	Last_name    string    `json:"last_name"`
	Address      string    `json:"address"`
	Phone        string    `json:"phone"`
	Payment_date int       `json:"payment_date"`
	Service_id   int       `json:"service_id"`
	Router_id    int       `json:"router_id"`
	CreatedAt    time.Time ` json:"created_at"`
	UpdatedAt    time.Time ` json:"updated_at"`
}

type Clients []*Client
