package controllers

import (
	"go-dom-parser/domain"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

//SetupController - initialize controller
func SetupController(storage *domain.Storage) *Controller {

	return &Controller{
		Storage: storage,
	}
}

// GetSymbols - get all loaded symbols
func (controller *Controller) GetSymbols(context *gin.Context) {
	requestQuery := SymbolsGetRequest{}
	//requestQuery.TargetPercent = make([]float64, 2)

	if context.ShouldBind(&requestQuery) == nil {
		// log.Println(postReq.Type)
	}

	symbols := make([]domain.Symbol, 0)

	for _, symbol := range controller.Storage.GetSymbolsKeys() {
		// if controller.Storage.SkipFilter(symbol) {
		// 	continue
		// }

		s := controller.Storage.GetSymbol(symbol)

		if s.Ratings["marketbeat"].TragetPercent >= requestQuery.TargetPercents[0] && s.Ratings["marketbeat"].TragetPercent <= requestQuery.TargetPercents[1] {
			symbols = append(symbols, s)
		}
	}

	sort.Sort(SortAscBySymbolId(symbols))

	// ... filter by page
	totalPages := (len(symbols) / requestQuery.PageSize) + 1

	from := (requestQuery.Page - 1) * requestQuery.PageSize
	if from > len(symbols) {
		from = 0
	}
	to := ((requestQuery.Page - 1) * requestQuery.PageSize) + requestQuery.PageSize
	if to > len(symbols) {
		to = len(symbols)
	}
	symbols = symbols[from:to]

	// ...

	if requestQuery.SortTargetPercent == "asc" {
		sort.Sort(SortAscByRatingPercent(symbols))

	}
	if requestQuery.SortTargetPercent == "desc" {
		sort.Sort(SortDescByRatingPercent(symbols))
	}

	response := SymbolsGetResponse{
		Symbols:    symbols,
		Filters:    controller.Storage.GetActiveFilterKeys(),
		TotalPages: totalPages,
	}

	response.BuildMarkers()

	context.JSON(http.StatusOK, response)
}

// DownloadResource - load resource by URI and send to parser
func (c *Controller) DownloadResource(context *gin.Context) {
	// postReq := ResourceRequest{}

	// if context.ShouldBind(&postReq) == nil {
	// 	// log.Println(postReq.Type)
	// }

	// // loading SPB Exchange Foreign Shares
	// if postReq.Type == "resource.spbexchange.foreignshares" {
	// 	core.SpbLoadForeignShares(c.RChan, postReq.Type)
	// }

	// // loading MarketBeat Ratings
	// if postReq.Type == "resource.marketbeat.ratings" {

	// 	for _, share := range c.Proc.SPBStocks {
	// 		core.MarketBeatLoadRatings(c.RChan, postReq.Type, share)
	// 	}
	// }

	// // loading Nasdaq Summary
	// if postReq.Type == "resource.nasdaq.summary" {
	// 	for _, share := range c.Proc.SPBStocks {
	// 		core.NasdaqLoadSummary(c.RChan, postReq.Type, share)
	// 	}
	// }

	context.JSON(http.StatusOK, "")
}
