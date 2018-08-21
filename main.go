package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"

	"goauth/config"
	"goauth/middlewares"
	"goauth/controllers"
	"go_todo_api/app"
)

func main() {
	ConfigRuntime()
	ConfigEnv()
	StartGin()
}

func ConfigRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Runing with %d CPUs\n", numCPU)
}

func ConfigEnv() {
	config.SetEnv()
}

func StartGin() {
	app.InitApp()
	defer app.CloseApp()

	router := gin.Default()

	router.Use(middlewares.ConnectDB(app.DBSession))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Auth API")
	})

	v1 := router.Group("v1")
	{
		v1.GET("/todos", controllers.IndexTodos)
		v1.POST("/todos", controllers.CreateTodo)
	}

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
