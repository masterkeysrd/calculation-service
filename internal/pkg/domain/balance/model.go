package balance

import "time"

type Balance struct {
	ID             uint    `json:"id"`
	UserID         uint    `json:"user_id"`
	Amount         float64 `json:"amount"`
	InFlightAmount float64 `json:"in_flight"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      time.Time
}

type NewBalanceInput struct {
	UserID uint
	Amount float64
}

func NewBalance(input NewBalanceInput) *Balance {
	return &Balance{
		UserID:         input.UserID,
		Amount:         input.Amount,
		InFlightAmount: 0,
	}
}

func (b *Balance) Reserve(amount float64) error {
	if b.Amount < amount+b.InFlightAmount {
		return ErrInsufficientFunds
	}

	b.InFlightAmount += amount
	return nil
}

func (b *Balance) Release(amount float64) error {
	if b.InFlightAmount < amount {
		return ErrInsufficientFunds
	}

	b.InFlightAmount -= amount
	return nil
}

func (b *Balance) Confirm(amount float64) error {
	if b.InFlightAmount < amount {
		return ErrInsufficientFunds
	}

	b.Amount -= amount
	b.InFlightAmount -= amount
	return nil
}

func (b *Balance) Rollback(amount float64) error {
	b.Amount += amount
	return nil
}
