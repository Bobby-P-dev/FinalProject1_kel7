package router

import (
	"github.com/Bobby-P-dev/FinalProject1_kel7/controllers"
	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todo", controllers.GetAllTodos)
	route.GET("/todo/:id", controllers.GetById)
	route.PUT("/todo/:id", controllers.PutById)
	route.DELETE("/todo/:id", controllers.DeleteTodo)

	route.Run()
}
