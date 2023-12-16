package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID           uuid.UUID       `json:"id,omitempty"`
	First_Name   string          `json:"first_name"`
	Last_name    string          `json:"last_name"`
	Address      string          `json:"address"`
	Phone        string          `json:"phone"`
	User_Data    json.RawMessage `json:"user_data"`
	Payment_date int             `json:"payment_date"`
	Service_id   uuid.UUID       `json:"service_id"`
	CreatedAt    time.Time       `json:"created_at"`
	UpdatedAt    time.Time       `json:"updated_at"`
}

type Clients []*Client
