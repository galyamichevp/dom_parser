package core

import (
	"encoding/json"
	"go-dom-parser/domain"
	"log"
	"sort"
	"time"
)

func (processor *Processor) nasdaqParseInfo(payload string) {
	var content struct {
		Symbol  string
		Content string
	}

	json.Unmarshal([]byte(payload), &content)

	info := &domain.NasdaqInfo{}

	err := json.Unmarshal([]byte(content.Content), &info)
	if err != nil {
		log.Printf("ERROR: fail unmarshal nasdaq info: %s", err.Error)
		return
	}

	nasdaqInfo := domain.Info{}

	nasdaqInfo.Symbol = content.Symbol
	nasdaqInfo.CompanyName = info.Data.CompanyName
	nasdaqInfo.StockType = info.Data.StockType
	nasdaqInfo.Exchange = info.Data.Exchange
	nasdaqInfo.Volume = info.Data.KeyStats.Volume.Value
	nasdaqInfo.PreviousClose = info.Data.KeyStats.PreviousClose.Value
	nasdaqInfo.OpenPrice = info.Data.KeyStats.OpenPrice.Value
	nasdaqInfo.LastSalePrice, _ = domain.FindPercentValue(info.Data.PrimaryData.LastSalePrice)
	nasdaqInfo.NetChange = info.Data.PrimaryData.NetChange
	nasdaqInfo.DeltaIndicator = info.Data.PrimaryData.DeltaIndicator

	if nasdaqInfo.DeltaIndicator == "up" {
		nasdaqInfo.PercentageChange, _ = domain.FindPercentValue(info.Data.PrimaryData.PercentageChange)
	}
	if nasdaqInfo.DeltaIndicator == "down" {
		nasdaqInfo.PercentageChange, _ = domain.FindPercentValue(info.Data.PrimaryData.PercentageChange)
		nasdaqInfo.PercentageChange *= -1
	}

	processor.Storage.SetInfo("nasdaq", nasdaqInfo)

	markerDelta := domain.Marker{
		Symbol: content.Symbol,
		FValue: nasdaqInfo.PercentageChange,
		BValue: nasdaqInfo.PercentageChange < -5,
	}

	processor.Storage.SetMarker("delta", markerDelta)

	log.Printf("INFO: nasdaq info loaded. Symbol: %s", content.Symbol)
}

func (processor *Processor) nasdaqParseSummary(payload string) {
	var content struct {
		Symbol  string
		Content string
	}

	json.Unmarshal([]byte(payload), &content)

	summary := &domain.NasdaqSummary{}

	err := json.Unmarshal([]byte(content.Content), &summary)
	if err != nil {
		log.Printf("ERROR: fail unmarshal nasdaq info: %s", err.Error)
		return
	}

	high, low, _ := domain.FindHighLowPriceValue(summary.Data.SummaryData.TodayHighLow.Value)
	high52, low52, _ := domain.FindHighLowPriceValue(summary.Data.SummaryData.FiftTwoWeekHighLow.Value)
	eps, _ := domain.FindPercentValue(summary.Data.SummaryData.EarningsPerShare.Value)
	prevClose, _ := domain.FindPercentValue(summary.Data.SummaryData.PreviousClose.Value)

	nasdaqSummary := domain.Summary{}

	nasdaqSummary.Symbol = content.Symbol
	nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	nasdaqSummary.Industry = summary.Data.SummaryData.Industry.Value
	nasdaqSummary.TodayHigh = high
	nasdaqSummary.TodayLow = low
	nasdaqSummary.ShareVolume = summary.Data.SummaryData.ShareVolume.Value
	nasdaqSummary.AverageVolume = summary.Data.SummaryData.AverageVolume.Value
	nasdaqSummary.PreviousClose = prevClose
	nasdaqSummary.FiftTwoWeekHigh = high52
	nasdaqSummary.FiftTwoWeekLow = low52
	nasdaqSummary.EarningsPerShare = eps

	processor.Storage.SetSummary("nasdaq", nasdaqSummary)

	markerEps := domain.Marker{
		Symbol: content.Symbol,
		FValue: eps,
		BValue: eps > 0,
	}

	processor.Storage.SetMarker("eps", markerEps)

	markerRange52 := domain.Marker{
		Symbol: content.Symbol,
		FValue: ((high52 - prevClose) / prevClose) * 100,
		BValue: ((prevClose - low52) / low52) < ((high52 - prevClose) / prevClose),
	}

	processor.Storage.SetMarker("range52", markerRange52)

	markerVolatility := domain.Marker{
		Symbol: content.Symbol,
		FValue: (((high - low) / low) * 100),
		BValue: (((high - low) / low) * 100) > 10,
	}

	processor.Storage.SetMarker("volatility", markerVolatility)

	log.Printf("INFO: nasdaq summary loaded. Symbol: %s", content.Symbol)
}

func (processor *Processor) nasdaqParseRealTime(payload string) {
	var content struct {
		Symbol  string
		Content []string
	}

	json.Unmarshal([]byte(payload), &content)

	var rtCollection []domain.NasdaqRealTime

	for _, item := range content.Content {
		rt := domain.NasdaqRealTime{}
		err := json.Unmarshal([]byte(item), &rt)
		if err != nil {
			log.Printf("ERROR: fail unmarshal nasdaq info: %s", err.Error)
			return
		}

		rtCollection = append(rtCollection, rt)
	}

	trades := make([]domain.Trade, 0)

	for _, item := range rtCollection {
		for _, row := range item.Data.Rows {
			trade := domain.Trade{}
			trade.Symbol = content.Symbol

			trade.Price, _ = domain.FindPriceValue(row.NLSPrice)
			trade.Volume = domain.ToFloat(row.NLSShareVolume)
			trade.Time = domain.ToTime(row.NLSTime)

			if trade.Time.Format("15:04:05") != "00:00:00" {
				trades = append(trades, trade)
			}
		}
	}

	sort.Sort(SortByDateTrade(trades))

	processor.Storage.SetTrades("nasdaq", content.Symbol, trades)

	// ...

	timeLimit, _ := time.Parse("15:04:05", trades[0].Time.Add(time.Hour*-2).Format("15:04:05"))

	var tradeIndex int

	for index, trade := range trades {
		if trade.Time.After(timeLimit) {
			tradeIndex = index
			continue
		}
		break
	}

	log.Printf("START TRADE")
	log.Println(trades[0])
	log.Println(trades[1])
	log.Println(trades[2])
	log.Printf("LIMIT TRADE")
	log.Println(trades[tradeIndex])
	log.Println(trades[tradeIndex+1])
	log.Println(trades[tradeIndex+2])

	markerDeviation := domain.Marker{
		Symbol: content.Symbol,
		FValue: ((trades[tradeIndex].Price - trades[0].Price) / trades[0].Price) * -100,
		BValue: (((trades[tradeIndex].Price - trades[0].Price) / trades[0].Price) * -100) > -2,
	}

	processor.Storage.SetMarker("deviation", markerDeviation)
	// ...

	log.Printf("INFO: nasdaq realtime loaded. Symbol: %s", content.Symbol)
}

func (processor *Processor) nasdaqParseHistory(payload string) {
	var content struct {
		Symbol  string
		Content string
	}

	json.Unmarshal([]byte(payload), &content)
	history := &domain.NasdaqHistory{}

	err := json.Unmarshal([]byte(content.Content), &history)
	if err != nil {
		log.Printf("ERROR: fail unmarshal nasdaq history: %s", err.Error)
		return
	}

	nasdaqHistory := domain.History{}

	nasdaqHistory.Symbol = content.Symbol
	for _, candle := range history.Data.Chart {
		chart := domain.Chart{
			High:     domain.ToFloat(candle.Z.High),
			Low:      domain.ToFloat(candle.Z.Low),
			Open:     domain.ToFloat(candle.Z.Open),
			Close:    domain.ToFloat(candle.Z.Close),
			Volume:   domain.ToFloat(candle.Z.Volume),
			DateTime: domain.ToDate(candle.Z.DateTime),
			Value:    domain.ToFloat(candle.Z.Value),
		}

		nasdaqHistory.Chart = append(nasdaqHistory.Chart, chart)
	}

	processor.Storage.SetHistory("nasdaq", nasdaqHistory)

	days1 := nasdaqHistory.Chart[4:5][0]
	days3 := nasdaqHistory.Chart[2:3][0]
	days5 := nasdaqHistory.Chart[:1][0]

	markerDelta3 := domain.Marker{
		Symbol: content.Symbol,
		FValue: ((days3.Close - days1.Close) / days1.Close) * -100,
		BValue: (((days3.Close - days1.Close) / days1.Close) * -100) < -5,
	}

	processor.Storage.SetMarker("delta3", markerDelta3)

	markerDelta5 := domain.Marker{
		Symbol: content.Symbol,
		FValue: ((days5.Close - days1.Close) / days1.Close) * -100,
		BValue: (((days5.Close - days1.Close) / days1.Close) * -100) < -10,
	}

	processor.Storage.SetMarker("delta5", markerDelta5)

	log.Printf("INFO: nasdaq history loaded. Symbol: %s", content.Symbol)
}
