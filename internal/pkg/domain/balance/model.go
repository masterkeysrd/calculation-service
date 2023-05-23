package balance

type Balance struct {
	ID             uint
	UserID         uint
	Amount         float64
	AmountInFlight float64
}

type NewBalanceInput struct {
	UserID uint
	Amount float64
}

func NewBalance(input NewBalanceInput) *Balance {
	return &Balance{
		UserID:         input.UserID,
		Amount:         input.Amount,
		AmountInFlight: 0,
	}
}

func (b *Balance) Reserve(amount float64) error {
	if b.Amount < amount+b.AmountInFlight {
		return ErrInsufficientFunds
	}

	b.AmountInFlight += amount
	return nil
}

func (b *Balance) Release(amount float64) error {
	if b.AmountInFlight < amount {
		return ErrInsufficientFunds
	}

	b.AmountInFlight -= amount
	return nil
}

func (b *Balance) Confirm(amount float64) error {
	if b.AmountInFlight < amount {
		return ErrInsufficientFunds
	}

	b.Amount -= amount
	b.AmountInFlight -= amount
	return nil
}

func (b *Balance) Rollback(amount float64) error {
	b.Amount += amount
	return nil
}
