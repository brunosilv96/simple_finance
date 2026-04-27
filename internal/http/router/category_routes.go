package router

import "github.com/gin-gonic/gin"

func registerCategoryRoutes(v1 *gin.RouterGroup, h CategoryHandler) {
	categories := v1.Group("/categories")
	{
		categories.POST("", h.Create)
		categories.GET("", h.FindAll)
		categories.GET("/:id", h.FindByID)
		categories.DELETE("/:id", h.Delete)
	}
}
