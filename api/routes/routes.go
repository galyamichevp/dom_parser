package routes

import (
	"go-dom-parser/api/controllers"

	"github.com/gin-gonic/gin"
)

// type Router gin.Engine

// func getItem(c *gin.Context) {
// 	c.String(200, "Success")
// }

// func (r *Router) InitRouters() {
// }

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("todo", controllers.GetTodos)
		// v1.POST("todo", Controllers.CreateATodo)
		// v1.GET("todo/:id", Controllers.GetATodo)
		// v1.PUT("todo/:id", Controllers.UpdateATodo)
		// v1.DELETE("todo/:id", Controllers.DeleteATodo)
	}
	return r
}
