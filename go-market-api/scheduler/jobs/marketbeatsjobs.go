package jobs

import (
	"encoding/json"
	"fmt"
	"go-dom-parser/domain"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MarketBeat struct {
	AnalyticsChan chan string
	Storage       *domain.Storage
}

//RunJob - run spb exchange job
func (marketBeat MarketBeat) RunJob() {
	go func() {
		for {
			if !marketBeat.Storage.GetSync("marketbeat.loadratings") {
				continue
			}

			marketBeat.Storage.SetSync("marketbeat.loadratings", false)

			for _, symbol := range marketBeat.Storage.GetSymbolsKeys() {

				if marketBeat.Storage.SkipFilter(symbol) {
					continue
				}

				go func(symbol string) {
					marketBeat.marketBeatLoadAnalytics(symbol)
				}(symbol)
			}
		}
	}()
}

// marketBeatLoadAnalytics - load MarketBeat analytics page
func (marketBeat MarketBeat) marketBeatLoadAnalytics(symbol string) {
	var content string
	content = marketbeatRequestResource(symbol, "NYSE")
	content = content + marketbeatRequestResource(symbol, "NASDAQ")

	var payload struct {
		Symbol  string
		Content string
	}

	payload.Symbol = symbol
	payload.Content = content

	out, err := json.Marshal(payload)
	if err != nil {
		log.Printf("ERROR: fail marshal: %s", err.Error)
	}

	marketBeat.AnalyticsChan <- string(out)
}

func marketbeatRequestResource(symbol string, market string) string {
	url := "https://www.marketbeat.com/stocks/" + market + "/" + symbol + "/price-target/"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "__cfduid=dc5e7c74b8f21b5e474a454d5481dd8281591791238; ASP.NET_SessionId=")

	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return ""
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return string(body)
}
