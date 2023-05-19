package operation

type Repository interface {
	FindAll() ([]Operation, error)
	FindByID(id uint64) (Operation, error)
	Create(operation Operation) error
	Update(operation Operation) error
	Delete(operation Operation) error
}
