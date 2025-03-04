package main

import (
	"log"
	"soap_service/internal/application"
	"time"
)

func main() {
	app := application.New()

	for {
		task, exists := app.Orch.GetTask()
		if !exists {
			log.Println("Нет доступных задач")
			time.Sleep(1 * time.Second)
			continue
		}

		calculator := app.Orch.GetCalculator()
		result, err := calculator.PerformOperation(task)
		if err != nil {
			log.Fatalf("Ошибка выполнения операции: %v", err)
		}

		err = app.Orch.SubmitTaskResult(task.ID, result)
		if err != nil {
			log.Fatalf("Не удалось отправить результат задачи: %v", err)
		}

		log.Printf("Задача %d выполнена с результатом: %f", task.ID, result)
	}

}
