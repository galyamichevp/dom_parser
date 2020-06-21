package routes

import (
	"go-dom-parser/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter - initialize routes
func SetupRouter(controller *controllers.Controller) *gin.Engine {

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

	r.Use(cors.New(config))

	v1 := r.Group("/api/v1")
	{
		v1.GET("symbols", controller.GetSymbols)
		v1.GET("resources/download", controller.DownloadResource)

		v1.GET("filters", controller.GetFilters)
		v1.POST("filters", controller.SetFilter)
	}
	return r
}
