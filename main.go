package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
)

func main() {
	ConfigRuntime()
	StartGin()
}

func ConfigRuntime() {
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	fmt.Printf("Runing with %d CPUs\n", numCPU)
}

func StartGin() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello Auth API")
	})

  v1 := router.Group("v1") {
    v1.GET("/todos", controllers.IndexTodos)
  }

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	router.Run(":" + port)
}
