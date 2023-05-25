package calculation

import (
	"math"
	"strconv"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/operation"
)

func addition(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments, 2)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]
	result := number1 + number2
	return formatNumber(result), nil
}

func subtraction(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments, 2)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]
	result := number1 - number2

	return formatNumber(result), nil
}

func multiplication(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments, 2)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]
	result := number1 * number2

	return formatNumber(result), nil
}

func division(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments, 2)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	number2 := numbers[1]
	result := number1 / number2

	return formatNumber(result), nil
}

func squareRoot(arguments []string) (string, error) {
	numbers, err := parseArguments(arguments, 1)
	if err != nil {
		return "", err
	}

	number1 := numbers[0]
	result := math.Sqrt(number1)

	return formatNumber(result), nil
}

func parseArguments(arguments []string, n int) ([]float64, error) {
	var numbers []float64

	if len(arguments) < n {
		return nil, ErrInvalidNumberOfArguments
	}

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	return numbers, nil
}

func performOperation(operationType operation.OperationType, arguments []string) (string, error) {
	switch operationType {
	case operation.OperationTypeAddition:
		return addition(arguments)
	case operation.OperationTypeSubtraction:
		return subtraction(arguments)
	case operation.OperationTypeMultiplication:
		return multiplication(arguments)
	case operation.OperationTypeDivision:
		return division(arguments)
	case operation.OperationTypeSquareRoot:
		return squareRoot(arguments)
	default:
		return "", ErrOperationNotSupported
	}
}

func formatNumber(number float64) string {
	return strconv.FormatFloat(number, 'f', -1, 64)
}
