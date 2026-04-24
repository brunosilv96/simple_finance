package app

import (
	http "github.com/brunosilv96/simple_finance_api/internal/http/handlers"
	"github.com/brunosilv96/simple_finance_api/internal/http/router"
	"github.com/brunosilv96/simple_finance_api/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

func App() *gin.Engine {
	// Infra
	categoryRepo := repository.NewMemoryCategory()
	userRepo := repository.NewMemoryUser()

	// Handlers with Use Cases
	categoryHandler := http.NewCategoryHandler(categoryRepo)
	userHandler := http.NewUserHandler(userRepo)

	return router.NewRouter(router.Router{
		HealthHandler:   http.HealthCheck,
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
	})
}
