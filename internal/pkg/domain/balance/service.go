package balance

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
