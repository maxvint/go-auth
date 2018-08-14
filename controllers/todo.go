package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"goauth/app"
	"goauth/models"
	"goauth/utils"
)

func IndexTodos(c *gin.Context) {
	c.JSON(http.StatusOK, "Index Todos")
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.BindJSON(&todo); err != nil {
		return
	}

	if err := todo.Create(app.GetDB(c)); err != nil {
		utils.AbortWithPublicError(c, http.StatusInternalServerError, err, "Couldn't create the todo")
		return
	}

	fmt.Println("323232323232")

	c.JSON(http.StatusCreated, todo)
}
