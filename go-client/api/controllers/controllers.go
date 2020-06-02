package controllers

import (
	"go-dom-parser/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

//deprecated
func GetTodos(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//GetResource - get resourse
func GetResource(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//GetResources - get all resourses
func GetResources(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//GetResourceResult - get resourse parsed result
func GetResourceResult(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//PostResourceDownload - load resource by URI and send to parser
func PostResourceDownload(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//PostResourceParse - send downloaded resource to parse
func PostResourceParse(c *gin.Context) {
	var todo []domain.Todo
	err := domain.GetAllTodos(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
