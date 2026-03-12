package api

import (
	"net/http"
	"simple_finance/internal/finance/handler"
)

func Router(categoryHandler *handler.CreateCategoryHandler) {
	http.HandleFunc("/categories", categoryHandler.CreateCategory)
}