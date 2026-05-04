package app

import (
	"log"

	http "github.com/brunosilv96/simple_finance_api/internal/http/handlers"
	"github.com/brunosilv96/simple_finance_api/internal/http/router"
	"github.com/brunosilv96/simple_finance_api/internal/infra"
	"github.com/brunosilv96/simple_finance_api/internal/infra/repository"
	"github.com/gin-gonic/gin"
)

func Run() *gin.Engine {
	// Infra
	_, err := infra.InitializePostgresDB()
	if err != nil {
		log.Fatal("error to open db connection")
	}

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
