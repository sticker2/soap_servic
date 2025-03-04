package main

import (
	"log"
	"net/http"
	"soap_service/internal/application"

	"github.com/gorilla/mux"
)

func main() {
	app := application.New()
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/calculate", app.AddExpression).Methods("POST")
	router.HandleFunc("/api/v1/expressions", app.GetExpressions).Methods("GET")
	router.HandleFunc("/api/v1/expressions/{id}", app.GetExpression).Methods("GET")

	router.HandleFunc("/internal/task", app.GetTask).Methods("GET")

	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
