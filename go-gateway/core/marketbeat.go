package core

import (
	"encoding/json"
	"fmt"
	"go-dom-parser/api/sockets"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// MarketBeatLoadRatings - load MarketBeat ratings
func MarketBeatLoadRatings(parserChan chan string, requestType string, share SPBStock) {
	// if item.Marker != "CHK" {
	// 	continue
	// }

	var content string
	if len(share.Marker) > 0 {
		content = marketbeatRequestResource(share, "NYSE")
		content = content + marketbeatRequestResource(share, "NASDAQ")
	}

	// ...
	pkg := sockets.OutPayload{
		Type:    requestType,
		Content: content,
		Marker:  share.Marker,
	}
	out, err := json.Marshal(pkg)
	if err != nil {
		log.Printf("ERROR: fail marshl: %s", err.Error)
	}

	parserChan <- string(out)
}

func marketbeatRequestResource(share SPBStock, market string) string {
	url := "https://www.marketbeat.com/stocks/" + market + "/" + share.Marker + "/price-target/"
	method := "GET"

	payload := strings.NewReader("")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Cookie", "__cfduid=dc5e7c74b8f21b5e474a454d5481dd8281591791238; ASP.NET_SessionId=")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		//fmt.Println(string(body))

		return string(body)
	}

	return ""
}
