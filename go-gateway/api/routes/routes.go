package routes

import (
	"go-dom-parser/api/controllers"

	"github.com/gin-gonic/gin"
)

//SetupRouter - initialize routes
func SetupRouter(ctrl controllers.Controller) *gin.Engine {

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("stocks", ctrl.GetStocks)
		v1.GET("resources/download", ctrl.DownloadResource)
		// v1.GET("resources/:id", controllers.GetResource)
		// v1.GET("resources/result/:id", controllers.GetResourceResult)
		// v1.POST("resources/parse/:id", controllers.PostResourceParse)
	}
	return r
}
