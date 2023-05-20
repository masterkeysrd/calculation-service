package operation

type Repository interface {
	FindAll() ([]Operation, error)
	FindByID(id uint64) (Operation, error)
}

var operations = []Operation{
	{
		ID:   1,
		Type: OperationTypeAddition,
		Cost: 0.5,
	},
	{
		ID:   2,
		Type: OperationTypeSubtraction,
		Cost: 0.5,
	},
	{
		ID:   3,
		Type: OperationTypeMultiplication,
		Cost: 0.5,
	},
	{
		ID:   4,
		Type: OperationTypeDivision,
		Cost: 0.5,
	},
	{
		ID:   5,
		Type: OperationTypeRandomString,
		Cost: 1.5,
	},
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindByID(id uint64) (Operation, error) {
	for _, operation := range operations {
		if operation.ID == id {
			return operation, nil
		}
	}

	return Operation{}, nil
}

func (r *repository) FindAll() ([]Operation, error) {
	return operations, nil
}
