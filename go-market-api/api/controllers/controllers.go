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
	// arr := make([]core.SPBStock, len(c.Proc.SPBStocks))

	// r, _ := regexp.Compile("[0-9,.]+")

	// for i := 0; i < len(c.Proc.SPBStocks); i++ {

	// 	if c.Proc.SPBStocks[i].Today != nil {
	// 		res := r.FindString(c.Proc.SPBStocks[i].Today[5])

	// 		res = strings.ReplaceAll(res, ",", "")

	// 		f, err := strconv.ParseFloat(res, 64)
	// 		if err == nil {
	// 			c.Proc.SPBStocks[i].Percent = f
	// 		}
	// 	}

	// 	arr[i] = c.Proc.SPBStocks[i]
	// }

	// sort.Sort(ByPercent(arr))

	requestQuery := SymbolsGetRequest{}

	if context.ShouldBind(&requestQuery) == nil {
		// log.Println(postReq.Type)
	}

	symbols := make([]domain.Symbol, 0)

	for _, symbol := range controller.Storage.GetSymbolsKeys() {
		if controller.Storage.SkipFilter(symbol) {
			continue
		}

		s := controller.Storage.GetSymbol(symbol)
		symbols = append(symbols, s)
	}

	sort.Sort(SortAscBySymbolId(symbols))

	if requestQuery.SortByPercent == "asc" {
		sort.Sort(SortAscByRatingPercent(symbols))

	}
	if requestQuery.SortByPercent == "desc" {
		sort.Sort(SortDescByRatingPercent(symbols))
	}

	if requestQuery.PercentLimit != "" {
		// lim, _ := strconv.Atoi(requestQuery.PercentLimit)
		// symbols = symbols[:lim]
	}

	context.JSON(http.StatusOK, symbols)
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
