package balance

import "errors"

var (
	balances = []Balance{
		{
			ID:     1,
			UserID: 1,
			Amount: 100,
		},
	}
)

type Repository interface {
	FindByUserID(userID uint) (*Balance, error)
	Create(balance *Balance) error
	Update(balance *Balance) error
	Delete(balance *Balance) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindByUserID(userID uint) (*Balance, error) {
	for _, balance := range balances {
		if balance.UserID == userID {
			return &balance, nil
		}
	}

	return nil, errors.New("balance not found")
}

func (r *repository) Create(balance *Balance) error {
	balances = append(balances, *balance)
	return nil
}

func (r *repository) Update(balance *Balance) error {
	for i, b := range balances {
		if b.ID == balance.ID {
			balances[i] = *balance
			return nil
		}
	}

	return errors.New("balance not found")
}

func (r *repository) Delete(balance *Balance) error {
	for i, b := range balances {
		if b.ID == balance.ID {
			balances = append(balances[:i], balances[i+1:]...)
			return nil
		}
	}

	return errors.New("balance not found")
}
