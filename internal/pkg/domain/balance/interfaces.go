package balance

type Service interface {
	FindByUserID(userID uint64) (Balance, error)
	Create(balance Balance) error
	Reserve(userID uint64, amount float64) error
	Release(userID uint64, amount float64) error
	Commit(userID uint64, amount float64) error
	Rollback(userID uint64, amount float64) error
}
