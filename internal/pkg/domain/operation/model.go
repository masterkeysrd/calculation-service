package operation

type Operation struct {
	ID   uint          `json:"id"`
	Type OperationType `json:"type"`
	Cost float64       `json:"cost"`
}

type OperationType string

const (
	OperationTypeAddition       OperationType = "addition"
	OperationTypeSubtraction    OperationType = "subtraction"
	OperationTypeMultiplication OperationType = "multiplication"
	OperationTypeDivision       OperationType = "division"
	OperationTypeSquareRoot     OperationType = "square_root"
	OperationTypeRandomString   OperationType = "random_string"
)
