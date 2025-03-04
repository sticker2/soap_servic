package orchestrator_test

import (
	"soap_service/internal/orchestrator"
	"soap_service/pkg/calculation"
	"testing"
)

func TestAddExpression(t *testing.T) {
	calculator := calculation.NewCalculator()
	orch := orchestrator.NewOrchestrator(calculator)
	task := calculation.Task{Arg1: "5", Arg2: "3", Operation: "add"}

	id, err := orch.AddExpression("123", task)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expr, err := orch.GetExpression(id)
	if err != nil {
		t.Fatalf("Expected expression, got error: %v", err)
	}
	if expr.Status != "pending" {
		t.Fatalf("Expected status 'pending', got %s", expr.Status)
	}
}

func TestGetExpressions(t *testing.T) {
	calculator := calculation.NewCalculator()
	orch := orchestrator.NewOrchestrator(calculator)

	orch.AddExpression("123", calculation.Task{Arg1: "5", Arg2: "3", Operation: "add"})

	expressions, err := orch.GetExpressions()
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if len(expressions) == 0 {
		t.Fatal("Expected at least one expression")
	}
}

func TestSubmitTaskResult(t *testing.T) {
	calculator := calculation.NewCalculator()
	orch := orchestrator.NewOrchestrator(calculator)
	id, err := orch.AddExpression("123", calculation.Task{Arg1: "5", Arg2: "3", Operation: "add"})
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	err = orch.SubmitTaskResult(id, 8.0)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expr, err := orch.GetExpression(id)
	if err != nil {
		t.Fatalf("Expected expression, got error: %v", err)
	}
	if *expr.Result != 8.0 {
		t.Fatalf("Expected result 8.0, got %f", *expr.Result)
	}
}
