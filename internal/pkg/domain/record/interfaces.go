package record

type Service interface {
	FindByUserID(userID uint64) ([]Record, error)
	Create(record Record) error
	Delete(record Record) error
}
