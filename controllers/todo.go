package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"goauth/app"
	"goauth/models"
)

func IndexTodos(c *gin.Context) {
	c.JSON(http.StatusOK, "Index Todos")
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	fmt.Println(todo)

	if err := c.BindJSON(&todo); err != nil {
		return
	}

	if err := todo.Create(app.GetDB(c)); err != nil {
		return
	}

	c.JSON(http.StatusCreated, todo)
}
