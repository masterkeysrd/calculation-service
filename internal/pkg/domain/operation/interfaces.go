package operation

type Service interface {
	FindAll() ([]Operation, error)
	FindByName(name string) (Operation, error)
	Create(operation Operation) error
	Update(operation Operation) error
	Delete(operation Operation) error
}
