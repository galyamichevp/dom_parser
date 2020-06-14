package controllers

import (
	"go-dom-parser/core"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// type PayloadRMQ struct {
// 	Content string
// 	Type    string
// 	Marker  string
// }

type PostReq struct {
	Curl string `form:"curl"`
	Type string `form:"type"`
	Verb string `form:"verb"`
}

// GetStocks - get all loaded stocks
func (c *Controller) GetStocks(context *gin.Context) {
	arr := make([]core.SPBStock, len(c.Proc.SPBStocks))

	r, _ := regexp.Compile("[0-9,.]+")

	for i := 0; i < len(c.Proc.SPBStocks); i++ {

		if c.Proc.SPBStocks[i].Today != nil {
			res := r.FindString(c.Proc.SPBStocks[i].Today[5])

			res = strings.ReplaceAll(res, ",", "")

			f, err := strconv.ParseFloat(res, 64)
			if err == nil {
				c.Proc.SPBStocks[i].Percent = f
			}
		}

		arr[i] = c.Proc.SPBStocks[i]
	}

	sort.Sort(ByPercent(arr))

	context.JSON(http.StatusOK, arr)
}

// DownloadResource - load resource by URI and send to parser
func (c *Controller) DownloadResource(context *gin.Context) {
	postReq := PostReq{}

	if context.ShouldBind(&postReq) == nil {
		// log.Println(postReq.Type)
	}

	// loading SPB Exchange Foreign Shares
	if postReq.Type == "resource.spbexchange.foreignshares" {
		core.SpbLoadForeignShares(c.RChan, postReq.Type)
	}

	// loading MarketBeat Ratings
	if postReq.Type == "resource.marketbeat.ratings" {

		for _, share := range c.Proc.SPBStocks {
			core.MarketBeatLoadRatings(c.RChan, postReq.Type, share)
		}
	}

	context.JSON(http.StatusOK, "")
}
