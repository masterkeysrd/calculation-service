package balance

type Service interface {
	FindByUserID(userID uint64) (Balance, error)
	Create(balance Balance) error
	Reserve(userID uint64, amount float64) error
	Release(userID uint64, amount float64) error
	Commit(userID uint64, amount float64) error
	Rollback(userID uint64, amount float64) error
}

type balanceService struct{}

func NewBalanceService() Service {
	return &balanceService{}
}

func (s *balanceService) FindByUserID(userID uint64) (Balance, error) {
	return Balance{}, nil
}

func (s *balanceService) Create(balance Balance) error {
	return nil
}

func (s *balanceService) Reserve(userID uint64, amount float64) error {
	return nil
}

func (s *balanceService) Release(userID uint64, amount float64) error {
	return nil
}

func (s *balanceService) Commit(userID uint64, amount float64) error {
	return nil
}

func (s *balanceService) Rollback(userID uint64, amount float64) error {
	return nil
}
