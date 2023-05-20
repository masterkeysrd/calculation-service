package calculation

import (
	"errors"
	"strconv"
)

func addition(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]

	return strconv.FormatFloat(number1+number2, 'f', -1, 64), nil
}

func subtraction(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]

	return strconv.FormatFloat(number1-number2, 'f', -1, 64), nil
}

func multiplication(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]

	return strconv.FormatFloat(number1*number2, 'f', -1, 64), nil
}

func division(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]

	return strconv.FormatFloat(number1/number2, 'f', -1, 64), nil
}

func parseArguments(arguments []string) ([]float64, error) {
	if len(arguments) < 2 {
		return nil, errors.New("invalid number of arguments, expected to have 2")
	}

	if len(arguments) > 2 {
		return nil, errors.New("invalid number of arguments, expected to have 2")
	}

	number1, err := strconv.ParseFloat(arguments[0], 64)
	if err != nil {
		return nil, err
	}

	number2, err := strconv.ParseFloat(arguments[1], 64)
	if err != nil {
		return nil, err
	}

	return []float64{number1, number2}, nil
}
