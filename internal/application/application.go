package application

import (
	"encoding/json"
	"net/http"
	"soap_service/internal/orchestrator"
	"soap_service/pkg/calculation"
	"strconv"

	"github.com/gorilla/mux"
)

type Application struct {
	Orch *orchestrator.Orchestrator
}

func New() *Application {
	calculator := calculation.NewCalculator()
	orch := orchestrator.NewOrchestrator(calculator)
	return &Application{Orch: orch}
}

func (app *Application) AddExpression(w http.ResponseWriter, r *http.Request) {
	if app.Orch == nil {
		http.Error(w, "orchestrator not initialized", http.StatusInternalServerError)
		return
	}

	var req struct {
		Expression string `json:"expression"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Expression == "" {
		http.Error(w, "invalid input", http.StatusUnprocessableEntity)
		return
	}

	task := calculation.Task{Arg1: req.Expression, Arg2: "", Operation: "add"}

	id, err := app.Orch.AddExpression(req.Expression, task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (app *Application) GetExpressions(w http.ResponseWriter, r *http.Request) {
	if app.Orch == nil {
		http.Error(w, "orchestrator not initialized", http.StatusInternalServerError)
		return
	}

	expressions, err := app.Orch.GetExpressions()
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"expressions": expressions})
}

func (app *Application) GetExpression(w http.ResponseWriter, r *http.Request) {
	if app.Orch == nil {
		http.Error(w, "orchestrator not initialized", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	expr, err := app.Orch.GetExpression(id)
	if err != nil {
		http.Error(w, "expression not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expr)
}

func (app *Application) GetTask(w http.ResponseWriter, r *http.Request) {
	if app.Orch == nil {
		http.Error(w, "orchestrator not initialized", http.StatusInternalServerError)
		return
	}

	task, exists := app.Orch.GetTask()
	if !exists {
		http.Error(w, "task not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{"task": task})
}

func (app *Application) PostTaskResult(w http.ResponseWriter, r *http.Request) {
	if app.Orch == nil {
		http.Error(w, "orchestrator not initialized", http.StatusInternalServerError)
		return
	}

	var req struct {
		Id     int     `json:"id"`
		Result float64 `json:"result"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid input", http.StatusUnprocessableEntity)
		return
	}

	if req.Result == 0 {
		http.Error(w, "invalid result", http.StatusUnprocessableEntity)
		return
	}

	err := app.Orch.SubmitTaskResult(req.Id, req.Result)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
