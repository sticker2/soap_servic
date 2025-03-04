package calculation

import "fmt"

func ErrorInvalidOperation(operation string) error {
	return fmt.Errorf("invalid operation: %s", operation)
}

func ErrorDivisionByZero() error {
	return fmt.Errorf("division by zero")
}
