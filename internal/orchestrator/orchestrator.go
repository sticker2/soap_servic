package orchestrator

import (
	"errors"
	"soap_service/pkg/calculation"
	"sync"
)

type Orchestrator struct {
	calculator  *calculation.Calculator
	mu          sync.Mutex
	expressions map[int]*Expression
	tasks       chan calculation.Task
}

type Expression struct {
	ID     int
	Status string
	Result *float64
}

func NewOrchestrator(calculator *calculation.Calculator) *Orchestrator {
	return &Orchestrator{
		calculator:  calculator,
		expressions: make(map[int]*Expression),
		tasks:       make(chan calculation.Task, 100),
	}
}

func (o *Orchestrator) GetCalculator() *calculation.Calculator {
	return o.calculator
}

func (o *Orchestrator) AddExpression(expression string, task calculation.Task) (int, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	if task.Arg1 == "" {
		return 0, errors.New("invalid task input")
	}

	id := len(o.expressions) + 1
	o.expressions[id] = &Expression{ID: id, Status: "pending"}
	o.tasks <- task
	return id, nil
}

func (o *Orchestrator) GetExpressions() ([]*Expression, error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	var expressions []*Expression
	for _, expr := range o.expressions {
		expressions = append(expressions, expr)
	}
	return expressions, nil
}

func (o *Orchestrator) GetExpression(id int) (*Expression, error) {
	o.mu.Lock()
	defer o.mu.Unlock()
	expr, exists := o.expressions[id]
	if !exists {
		return nil, errors.New("expression not found")
	}
	return expr, nil
}

func (o *Orchestrator) GetTask() (calculation.Task, bool) {
	o.mu.Lock()
	defer o.mu.Unlock()

	select {
	case task := <-o.tasks:
		return task, true
	default:
		return calculation.Task{}, false
	}
}

func (o *Orchestrator) SubmitTaskResult(id int, result float64) error {
	o.mu.Lock()
	defer o.mu.Unlock()
	expr, exists := o.expressions[id]
	if !exists {
		return errors.New("task not found")
	}
	expr.Status = "completed"
	expr.Result = &result
	return nil
}
