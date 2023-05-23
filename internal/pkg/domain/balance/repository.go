package balance

type Repository interface {
	GetWithUserID(userID uint) (*Balance, error)
	PerformTransaction(userID uint, transaction func(balance *Balance) error) (*Balance, error)
	Delete(balance *Balance) error
}
