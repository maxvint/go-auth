package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	status  int    `json: "status"`
	message string `json: "message"`
	data    string `json: "data"`
}

// func SuccessResponse() Response {
// 	return Response{Status: http.StatusCreated, Message: http.StatusText(http.StatusCreated)}
// }

func SuccessCreated(c *gin.Context, obj gin.H) {
	c.JSON(http.StatusCreated, gin.H(obj))
}

func SuccessOK(c *gin.Context, obj gin.H) {
	c.JSON(http.StatusOK, gin.H(obj))
}

func FailedNotFound(c *gin.Context, obj gin.H) {
	c.JSON(http.StatusNotFound, gin.H(obj))
}
