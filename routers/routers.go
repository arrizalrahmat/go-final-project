package routers

import (
	"go-final-project/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	todos := router.Group("/todos")
	{
		todos.GET("/", controllers.GetTodos)
		todos.GET("/:id", controllers.GetTodo)
		todos.POST("/", controllers.CreateTodo)
		todos.PUT("/:id", controllers.UpdateTodo)
		todos.DELETE("/:id", controllers.DeleteTodo)
	}

	return router
}
