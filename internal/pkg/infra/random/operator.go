package random

type Operator interface {
	GetMethod() string
	GetParams() map[string]interface{}
}
