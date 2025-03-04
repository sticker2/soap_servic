package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddExpression_Success(t *testing.T) {
	app := New()
	body, _ := json.Marshal(map[string]string{"expression": "2+2*2"})
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AddExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code 201, got %v", rr.Code)
	}
}

func TestAddExpression_InvalidExpression(t *testing.T) {
	app := New()
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AddExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code 422, got %v", rr.Code)
	}
}

func TestAddExpression_InternalError(t *testing.T) {
	app := New()
	req, err := http.NewRequest("POST", "/api/v1/calculate", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.AddExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %v", rr.Code)
	}
}

func TestGetExpressions_Success(t *testing.T) {
	app := New()
	req, err := http.NewRequest("GET", "/api/v1/expressions", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetExpressions)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rr.Code)
	}
}

func TestGetExpression_Success(t *testing.T) {
	app := New()
	req, err := http.NewRequest("GET", "/api/v1/expressions/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rr.Code)
	}
}

func TestGetExpression_NotFound(t *testing.T) {
	app := New()
	req, err := http.NewRequest("GET", "/api/v1/expressions/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", rr.Code)
	}
}

func TestGetExpression_InternalError(t *testing.T) {
	app := New()
	req, err := http.NewRequest("GET", "/api/v1/expressions/error", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetExpression)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code 500, got %v", rr.Code)
	}
}

func TestGetTask_NotFound(t *testing.T) {
	app := New()
	req, err := http.NewRequest("GET", "/internal/task", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.GetTask)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code 404, got %v", rr.Code)
	}
}

func TestPostTaskResult_InvalidResult(t *testing.T) {
	app := New()
	req, err := http.NewRequest("POST", "/internal/task/result", bytes.NewBuffer([]byte(`{}`)))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.PostTaskResult)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected status code 422, got %v", rr.Code)
	}
}

func TestPostTaskResult_Success(t *testing.T) {
	app := New()
	body, _ := json.Marshal(map[string]interface{}{
		"id":     1,
		"result": 8.5,
	})
	req, err := http.NewRequest("POST", "/internal/task/result", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(app.PostTaskResult)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code 200, got %v", rr.Code)
	}
}
