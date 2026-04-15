package main

import (
	"log"
	"net/http"

	"github.com/brunosilv96/simple_finance_api/internal/finance/handler"
	"github.com/brunosilv96/simple_finance_api/internal/finance/repository"
	"github.com/brunosilv96/simple_finance_api/internal/middleware"
)

func main() {
	// Dependency Injection
	// 1. Infra
	memoryCategoryRepository := repository.NewMemoryCategory()
	memoryUserRepository := repository.NewMemoryUser()

	// 3. Handlers
	categoryHandler := handler.NewCategoryHandler(memoryCategoryRepository)
	userHandler := handler.NewUserHandler(memoryUserRepository)

	// 4. Server
	mux := http.NewServeMux()

	// GET: /health
	mux.HandleFunc("GET /api/v1/health", handler.HealthCheck)

	// Setup Handlers
	categoryHandler.SetupRoutes(mux)
	userHandler.SetupRoutes(mux)

	// Start web server
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", middleware.Logger(mux))
	if err != nil {
		log.Fatal(err)
	}
}
