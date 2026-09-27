package calculator

import "errors"

var (
	ErrInvalidOperation = errors.New("Invalid operation")
	ErrDivisionByZero   = errors.New("Division by zero")
)

func Calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, ErrDivisionByZero
		}
		return a / b, nil
	default:
		return 0, ErrInvalidOperation
	}
}
