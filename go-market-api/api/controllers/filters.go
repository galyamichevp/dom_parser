package controllers

import (
	"go-dom-parser/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFilters - ...
func (controller *Controller) GetFilters(context *gin.Context) {

	filters := make([]domain.Filter, 0)
	for _, filter := range controller.Storage.Filters {
		filters = append(filters, filter)
	}

	context.JSON(http.StatusOK, filters)
}

// SetFilter - ...
func (controller *Controller) SetFilter(context *gin.Context) {

	requestQuery := FilterPostRequest{}
	context.BindJSON(&requestQuery)

	for _, symbol := range controller.Storage.Filters {
		filter := controller.Storage.Filters[symbol.Symbol]
		filter.IsActive = false
		controller.Storage.Filters[symbol.Symbol] = filter
	}

	for _, symbol := range requestQuery.Symbols {
		controller.Storage.Filters[symbol] = domain.Filter{
			IsActive: true,
			Symbol:   symbol,
		}
	}
}

// DeleteFilter - ...
func (controller *Controller) DeleteFilter(context *gin.Context) {
}
