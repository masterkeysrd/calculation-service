package balance

type BalanceService struct{}

func NewBalanceService() Service {
	return &BalanceService{}
}

func (s *BalanceService) FindByUserID(userID uint64) (Balance, error) {
	return Balance{}, nil
}

func (s *BalanceService) Create(balance Balance) error {
	return nil
}

func (s *BalanceService) Reserve(userID uint64, amount float64) error {
	return nil
}

func (s *BalanceService) Release(userID uint64, amount float64) error {
	return nil
}

func (s *BalanceService) Commit(userID uint64, amount float64) error {
	return nil
}

func (s *BalanceService) Rollback(userID uint64, amount float64) error {
	return nil
}
