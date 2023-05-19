package balance

type Repository interface {
	FindByUserID(userID uint64) (Balance, error)
	Create(balance Balance) error
	Update(balance Balance) error
	Delete(balance Balance) error
}
