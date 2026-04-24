package router

import "github.com/gin-gonic/gin"

type UserHandler interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
}

type CategoryHandler interface {
	Create(c *gin.Context)
	FindAll(c *gin.Context)
	FindByID(c *gin.Context)
	Delete(c *gin.Context)
}

type Router struct {
	HealthHandler   gin.HandlerFunc
	UserHandler     UserHandler
	CategoryHandler CategoryHandler
}

func NewRouter(router Router) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/health", router.HealthHandler)

	v1 := r.Group("/api/v1")
	registerUserRoutes(v1, router.UserHandler)
	registerCategoryRoutes(v1, router.CategoryHandler)

	return r
}
