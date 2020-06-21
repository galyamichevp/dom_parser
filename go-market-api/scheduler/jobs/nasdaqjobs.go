package jobs

import (
	"encoding/json"
	"fmt"
	"go-dom-parser/domain"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Nasdaq struct {
	InfoChan     chan string
	SummaryChan  chan string
	RealTimeChan chan string
	HistoryChan  chan string
	Storage      *domain.Storage
}

//RunJob - run nasdaq job
func (nasdaq Nasdaq) RunJob() {
	// time.Sleep(5 * time.Second)
	go func() {
		for {

			for _, symbol := range nasdaq.Storage.GetSymbolsKeys() {

				if nasdaq.Storage.SkipFilter(symbol) {
					continue
				}

				time.Sleep(300 * time.Millisecond)

				fmt.Println("INFO: nasdaq ... " + symbol)

				nasdaq.nasdaqLoadInfo(symbol)
				nasdaq.nasdaqLoadSummary(symbol)
				nasdaq.nasdaqLoadRealTime(symbol)
				nasdaq.nasdaqLoadHistory(symbol)
			}
		}
	}()
}

// nasdaqLoadInfo - load Nasdaq Info page
func (nasdaq Nasdaq) nasdaqLoadInfo(symbol string) {
	// if symbol != "RCL" {
	// 	return
	// }

	content := nasdaqRequestInfo(symbol)

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

	nasdaq.InfoChan <- string(out)
}

func nasdaqRequestInfo(symbol string) string {
	url := "https://api.nasdaq.com/api/quote/" + symbol + "/info?assetclass=stocks"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "api.nasdaq.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36")
	req.Header.Add("origin", "https://www.nasdaq.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.nasdaq.com/market-activity/stocks/momo")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		//fmt.Println(string(body))

		return string(body)
	}

	return ""
}

// nasdaqLoadSummary - load Nasdaq Summary page
func (nasdaq Nasdaq) nasdaqLoadSummary(symbol string) {
	// if symbol != "RCL" {
	// 	return
	// }

	content := nasdaqRequestSummary(symbol)

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

	nasdaq.SummaryChan <- string(out)
}

func nasdaqRequestSummary(symbol string) string {
	url := "https://api.nasdaq.com/api/quote/" + symbol + "/summary?assetclass=stocks"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "api.nasdaq.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	req.Header.Add("origin", "https://www.nasdaq.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.nasdaq.com/market-activity/stocks/anab")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cookie", "NSC_W.OEBR.DPN.7070=ffffffffc3a0f70e45525d5f4f58455e445a4a422dae")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		//fmt.Println(string(body))

		return string(body)
	}

	return ""
}

// nasdaqLoadRealTime - load Nasdaq RealTime data
func (nasdaq Nasdaq) nasdaqLoadRealTime(symbol string) {
	// if symbol != "RCL" {
	// 	return
	// }

	var content []string

	mmStartPoint := 30

	for hTimePoint := 9; hTimePoint < 16; hTimePoint++ {
		for mTimePoint := mmStartPoint; mTimePoint < 31; mTimePoint += 30 {

			hh := fmt.Sprintf("%02d", hTimePoint)
			mm := fmt.Sprintf("%02d", mTimePoint)

			content = append(content, nasdaqRequestRealTime(symbol, hh, mm))
		}

		mmStartPoint = 0
	}

	var payload struct {
		Symbol  string
		Content []string
	}

	payload.Symbol = symbol
	payload.Content = content

	out, err := json.Marshal(payload)
	if err != nil {
		log.Printf("ERROR: fail marshal: %s", err.Error)
	}

	nasdaq.RealTimeChan <- string(out)
}

func nasdaqRequestRealTime(symbol, hh, mm string) string {
	url := "https://api.nasdaq.com/api/quote/" + symbol + "/realtime-trades?&limit=50000&offset=0&fromTime=" + hh + ":" + mm
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "api.nasdaq.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	req.Header.Add("origin", "https://www.nasdaq.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.nasdaq.com/market-activity/stocks/rcl/latest-real-time-trades")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cookie", "NSC_W.OEBR.DPN.7070=ffffffffc3a08e0e45525d5f4f58455e445a4a422dae")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		//fmt.Println(string(body))

		return string(body)
	}

	return ""
}

// nasdaqLoadHistory - load Nasdaq History page
func (nasdaq Nasdaq) nasdaqLoadHistory(symbol string) {
	content := nasdaqRequestHistory(symbol)

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

	nasdaq.HistoryChan <- string(out)
}

func nasdaqRequestHistory(symbol string) string {
	currentTime := time.Now().Local()
	cDate := currentTime.Format("2006-01-02")
	pDate := currentTime.AddDate(0, 0, -7).Format("2006-01-02")

	url := "https://api.nasdaq.com/api/quote/" + symbol + "/chart?assetclass=stocks&fromdate=" + pDate + "&todate=" + cDate
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "api.nasdaq.com")
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("accept", "application/json, text/plain, */*")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36")
	req.Header.Add("origin", "https://www.nasdaq.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://www.nasdaq.com/market-activity/stocks/chk")
	req.Header.Add("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Add("Cookie", "NSC_W.OEBR.DPN.7070=ffffffffc3a0f73145525d5f4f58455e445a4a422dae")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode == http.StatusOK {
		//fmt.Println(string(body))

		return string(body)
	}

	return ""
}
