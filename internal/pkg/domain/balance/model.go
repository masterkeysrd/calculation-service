package balance

import "time"

type Balance struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount"`
	InFlight  float64 `json:"in_flight"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type NewBalanceInput struct {
	UserID uint
	Amount float64
}

func NewBalance(input NewBalanceInput) *Balance {
	return &Balance{
		UserID:   input.UserID,
		Amount:   input.Amount,
		InFlight: 0,
	}
}

func (b *Balance) Reserve(amount float64) error {
	if b.Amount < amount+b.InFlight {
		return ErrInsufficientFunds
	}

	b.InFlight += amount
	return nil
}

func (b *Balance) Release(amount float64) error {
	if b.InFlight < amount {
		return ErrInsufficientFunds
	}

	b.InFlight -= amount
	return nil
}

func (b *Balance) Confirm(amount float64) error {
	if b.InFlight < amount {
		return ErrInsufficientFunds
	}

	b.Amount -= amount
	b.InFlight -= amount
	return nil
}

func (b *Balance) Rollback(amount float64) error {
	b.Amount += amount
	return nil
}
