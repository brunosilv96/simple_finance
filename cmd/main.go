package main

import (
	"log"
	"net/http"
	"simple_finance/internal/api"
	"simple_finance/internal/finance/handler"
	"simple_finance/internal/finance/usecase"
)

func main() {
	// Dependency Injection
	createCategoryUseCase := usecase.CreateCategoryUseCase{}
	categoryHandler := handler.CreateCategoryHandler{
		CreateCategoryUseCase: createCategoryUseCase,
	}

	// GET: /health
	http.HandleFunc("/health", handler.HealthHandler)

	api.Router(&categoryHandler)

	
	// Start web server
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}