package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bobby-P-dev/FinalProject1_kel7/initializers"
	"github.com/Bobby-P-dev/FinalProject1_kel7/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo struct {
		Name        string
		Description string
	}

	c.ShouldBindJSON(&todo)
	todos := models.Todo{Name: todo.Name, Description: todo.Description}
	result := initializers.DB.Create(&todos)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "succes create data",
		"todos":   todos,
	})

}

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo

	err := initializers.DB.Find(&todos)
	if err.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success Get data",
		"data":    todos,
	})

}

func GetById(c *gin.Context) {
	db := initializers.GetDB()
	todoId, _ := strconv.Atoi(c.Param("id"))

	var todos []models.Todo
	Todos := models.Todo{}

	Todos.ID = uint(todoId)

	err := db.First(&Todos, todoId).Error

	if err != nil {
		return
	}

	err = db.Model(&Todos).Where("id = ?", todoId).Find(&todos, Todos).Error

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Todos,
	})

}

func DeleteTodo(c *gin.Context) {
	db := initializers.GetDB()
	todoId, _ := strconv.Atoi(c.Param("id"))
	Todos := models.Todo{}

	Todos.ID = uint(todoId)

	err := db.First(&Todos, todoId).Error

	if err != nil {
		return
	}
	db.Model(&Todos).Delete(&Todos)

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success delete data",
	})
}

func PutById(c *gin.Context) {
	db := initializers.GetDB()
	todoId, _ := strconv.Atoi(c.Param("id"))
	Todos := models.Todo{}

	Todos.ID = uint(todoId)

	var updateTodo struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	c.Bind(&updateTodo)

	err := db.First(&Todos, todoId).Error

	if err != nil {
		return
	}

	err = db.Model(&Todos).Where("id = ?", todoId).Updates(&models.Todo{
		Name: updateTodo.Name, Description: updateTodo.Description,
	}).Error

	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Todos,
	})

}
