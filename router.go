package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuwenhui/goauth/controllers"
)

func CreateRouter(r *gin.Engine) {
	r.GET("api/todo", controllers.IndexTodos)
}
