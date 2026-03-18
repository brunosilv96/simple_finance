package main

import (
	"log"
	"net/http"
	"simple_finance/internal/finance/handler"
	"simple_finance/internal/finance/repository"
	usecase "simple_finance/internal/finance/usecase/category"
	"simple_finance/internal/middleware"
)

func main() {
	// Dependency Injection
	// 1. Infra
	memoryCategoryRepository := repository.NewMemoryCategory()
	
	// 2. Business
	createCategoryUseCase := usecase.NewCreateCategory(memoryCategoryRepository)
	listCategoriesUseCase := usecase.NewListCategories(memoryCategoryRepository)
	categoryByIdUseCase := usecase.NewFindCategoryById(memoryCategoryRepository)

	// 3. Handlers
	categoryHandler := handler.NewCategoryHandler(*createCategoryUseCase, *listCategoriesUseCase, *categoryByIdUseCase)

	// 4. Server
	mux := http.NewServeMux()

	// GET: /health
	mux.HandleFunc("GET /api/v1/health", handler.HealthCheck)

	// Setup Handlers
	categoryHandler.SetupRoutes(mux)
	
	// Start web server
	log.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", middleware.Logger(mux))
	if err != nil {
		log.Fatal(err)
	}
}