package router

import "github.com/gin-gonic/gin"

func registerUserRoutes(v1 *gin.RouterGroup, h UserHandler) {
	users := v1.Group("/users")
	{
		users.POST("", h.Create)
		users.GET("/:id", h.FindByID)
	}
}
