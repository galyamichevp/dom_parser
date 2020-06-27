package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetSync - ...
func (controller *Controller) SetSync(context *gin.Context) {
	requestQuery := SyncPostRequest{}
	context.BindJSON(&requestQuery)

	controller.Storage.SetSync(requestQuery.Source, requestQuery.State)

	context.JSON(http.StatusOK, "")
}
