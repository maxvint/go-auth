package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"goauth/app"
	"goauth/models"
	"log"
)

func IndexTodos(c *gin.Context) {
	c.JSON(http.StatusOK, "Index Todos")
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		log.Fatalln("Todo bind json error", err)
		return
	}

	if err := todo.Create(app.GetDB(c)); err != nil {
		log.Fatalln("Create Todo Model err", err)
		return
	}

	c.JSON(http.StatusCreated, todo)
}
