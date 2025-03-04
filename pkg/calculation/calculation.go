package calculation

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct{}

type Task struct {
	ID        int
	Arg1      string
	Arg2      string
	Operation string
}

type Expression struct {
	Task   Task
	Status string
	Result *float64
}

func NewCalculator() *Calculator {
	return &Calculator{}
}

func (c *Calculator) PerformOperation(task Task) (float64, error) {
	task.Operation = strings.TrimSpace(strings.ToLower(task.Operation))

	arg1, err := strconv.ParseFloat(task.Arg1, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid argument: %s", task.Arg1)
	}

	arg2, err := strconv.ParseFloat(task.Arg2, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid argument: %s", task.Arg2)
	}

	switch task.Operation {
	case "add":
		return c.Add(arg1, arg2)
	case "subtract":
		return c.Subtract(arg1, arg2)
	case "multiply":
		return c.Multiply(arg1, arg2)
	case "divide":
		return c.Divide(arg1, arg2)
	default:
		return 0, fmt.Errorf("unsupported operation: %s", task.Operation)
	}
}

func (c *Calculator) Add(arg1, arg2 float64) (float64, error) {
	return arg1 + arg2, nil
}

func (c *Calculator) Subtract(arg1, arg2 float64) (float64, error) {
	return arg1 - arg2, nil
}

func (c *Calculator) Multiply(arg1, arg2 float64) (float64, error) {
	return arg1 * arg2, nil
}

func (c *Calculator) Divide(arg1, arg2 float64) (float64, error) {
	if arg2 == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return arg1 / arg2, nil
}
