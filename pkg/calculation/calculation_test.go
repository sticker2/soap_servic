package calculation

import "testing"

func TestCalculator_Add(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "5", Arg2: "3", Operation: "add"}

	result, err := calculator.PerformOperation(task)
	if err != nil {
		t.Fatalf("Error during calculation: %v", err)
	}

	expectedResult := 8.0
	if result != expectedResult {
		t.Fatalf("expected %f, got %f", expectedResult, result)
	}
}

func TestCalculator_Subtract(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "5", Arg2: "3", Operation: "subtract"}

	result, err := calculator.PerformOperation(task)
	if err != nil {
		t.Fatalf("Error during calculation: %v", err)
	}

	expectedResult := 2.0
	if result != expectedResult {
		t.Fatalf("expected %f, got %f", expectedResult, result)
	}
}

func TestCalculator_Multiply(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "5", Arg2: "3", Operation: "multiply"}

	result, err := calculator.PerformOperation(task)
	if err != nil {
		t.Fatalf("Error during calculation: %v", err)
	}

	expectedResult := 15.0
	if result != expectedResult {
		t.Fatalf("expected %f, got %f", expectedResult, result)
	}
}

func TestCalculator_Divide(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "6", Arg2: "3", Operation: "divide"}

	result, err := calculator.PerformOperation(task)
	if err != nil {
		t.Fatalf("Error during calculation: %v", err)
	}

	expectedResult := 2.0
	if result != expectedResult {
		t.Fatalf("expected %f, got %f", expectedResult, result)
	}
}

func TestCalculator_DivideByZero(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "5", Arg2: "0", Operation: "divide"}

	_, err := calculator.PerformOperation(task)
	if err == nil {
		t.Fatal("Expected error for division by zero")
	}
}

func TestCalculator_InvalidOperation(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "5", Arg2: "3", Operation: "invalid"}

	_, err := calculator.PerformOperation(task)
	if err == nil {
		t.Fatal("Expected error for unsupported operation")
	}
}

func TestCalculator_InvalidArgument(t *testing.T) {
	calculator := NewCalculator()
	task := Task{Arg1: "five", Arg2: "3", Operation: "add"}

	_, err := calculator.PerformOperation(task)
	if err == nil {
		t.Fatal("Expected error for invalid argument")
	}
}
