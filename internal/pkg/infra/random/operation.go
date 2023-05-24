package random

type operation struct {
	Method string
	Params map[string]interface{}
}

func (o *operation) GetMethod() string {
	return o.Method
}

func (o *operation) GetParams() map[string]interface{} {
	return o.Params
}
