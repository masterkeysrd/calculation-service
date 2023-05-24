package random

import "fmt"

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewOperation(method string, params map[string]interface{}) Operator {
	return &operation{
		Method: method,
		Params: params,
	}
}

func NewGenerateStringOperation(length int) Operator {
	if length < 8 {
		length = 8
	}

	if length > 32 {
		length = 32
	}

	return NewOperation("generateStrings", map[string]interface{}{
		"n":          "1",
		"length":     fmt.Sprintf("%d", length),
		"characters": characters,
	})
}
