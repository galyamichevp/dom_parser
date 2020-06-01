package controllers

import (
	"go-dom-parser/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

//List all todos
func GetTodos(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
