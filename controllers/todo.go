package controllers

import (
	"code.bigogo.com/gopkg/errno"
	"github.com/gin-gonic/gin"
	"github.com/yuwenhui/goauth/dal"
	"net/http"
)

func IndexTodos(c *gin.Context) (interface{}, *errno.ErrNo) {
	c.JSON(http.StatusOK, "Index Todo")
	todos, err := dal.GetTodos()

	if err != nil {
		return nil, errno.DBError
	}

	return todos, nil
}
