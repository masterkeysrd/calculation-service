package operation

type Operation struct {
	ID   uint64        `json:"id"`
	Type OperationType `json:"type"`
	Cost float64       `json:"cost"`
}

type OperationType string

const (
	OperationTypeAddition       OperationType = "addition"
	OperationTypeSubtraction    OperationType = "subtraction"
	OperationTypeMultiplication OperationType = "multiplication"
	OperationTypeDivision       OperationType = "division"
	OperationTypeRandomString   OperationType = "random_string"
)
