package record

type Repository interface {
	GetWithUserID(userID uint, id uint) (*Record, error)
	ListWithUserID(userID uint) ([]*Record, error)
	Create(record *Record) error
	Delete(record *Record) error
}
