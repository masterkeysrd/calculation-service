package record

type Repository interface {
	FindByUserID(userID uint64) ([]Record, error)
	FindByUserIDAndID(userID uint64, id uint64) (Record, error)
	Create(record Record) error
	Update(record Record) error
	Delete(record Record) error
}
