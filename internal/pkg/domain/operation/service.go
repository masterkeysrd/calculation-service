package operation

type Service interface {
	FindAll() ([]Operation, error)
	FindByName(name string) (Operation, error)
	Create(operation Operation) error
	Update(operation Operation) error
	Delete(operation Operation) error
}

type operationService struct{}

func NewOperationService() Service {
	return &operationService{}
}

func (s *operationService) FindAll() ([]Operation, error) {
	return nil, nil
}

func (s *operationService) FindByName(name string) (Operation, error) {
	return Operation{}, nil
}

func (s *operationService) Create(operation Operation) error {
	return nil
}

func (s *operationService) Update(operation Operation) error {
	return nil
}

func (s *operationService) Delete(operation Operation) error {
	return nil
}
