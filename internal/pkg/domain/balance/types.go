package balance

import "time"

type Balance struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	Amount    float64   `json:"amount"`
	InFlight  float64   `json:"in_flight"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
