package core

import (
	"encoding/json"
	"go-dom-parser/domain"
	"log"
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

	nasdaqSummary := domain.Summary{}

	nasdaqSummary.Symbol = content.Symbol
	nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	nasdaqSummary.Industry = summary.Data.SummaryData.Industry.Value
	nasdaqSummary.TodayHigh = high
	nasdaqSummary.TodayLow = low
	nasdaqSummary.TodayVolatility = (high - low) / low
	nasdaqSummary.ShareVolume = summary.Data.SummaryData.ShareVolume.Value
	nasdaqSummary.AverageVolume = summary.Data.SummaryData.AverageVolume.Value
	nasdaqSummary.PreviousClose = summary.Data.SummaryData.PreviousClose.Value
	nasdaqSummary.FiftTwoWeekHighLow = summary.Data.SummaryData.FiftTwoWeekHighLow.Value
	nasdaqSummary.EarningsPerShare, _ = domain.FindPercentValue(summary.Data.SummaryData.EarningsPerShare.Value)

	processor.Storage.SetSummary("nasdaq", nasdaqSummary)

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

	// realTimeColletion := make([]domain.NasdaqRealTime, len(content.Content))

	// for _, item := range content.Content {
	// 	realTime := domain.NasdaqRealTime{}

	// 	err := json.Unmarshal([]byte(item), &realTime)
	// 	if err != nil {
	// 		log.Printf("ERROR: fail unmarshal nasdaq info: %s", err.Error)
	// 		return
	// 	}

	// 	realTimeColletion = append(realTimeColletion, realTime)
	// }

	// nasdaqSummary := domain.Summary{}

	// nasdaqSummary.Symbol = content.Symbol
	// nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	// nasdaqSummary.Sector = summary.Data.SummaryData.Sector.Value
	// nasdaqSummary.Industry = summary.Data.SummaryData.Industry.Value
	// nasdaqSummary.TodayHighLow = summary.Data.SummaryData.TodayHighLow.Value
	// nasdaqSummary.ShareVolume = summary.Data.SummaryData.ShareVolume.Value
	// nasdaqSummary.AverageVolume = summary.Data.SummaryData.AverageVolume.Value
	// nasdaqSummary.PreviousClose = summary.Data.SummaryData.PreviousClose.Value
	// nasdaqSummary.FiftTwoWeekHighLow = summary.Data.SummaryData.FiftTwoWeekHighLow.Value
	// nasdaqSummary.EarningsPerShare, _ = domain.FindPercentValue(summary.Data.SummaryData.EarningsPerShare.Value)

	// processor.Storage.SetSummary("nasdaq", nasdaqSummary)

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
			DateTime: domain.ToTime(candle.Z.DateTime),
			Value:    domain.ToFloat(candle.Z.Value),
		}

		nasdaqHistory.Chart = append(nasdaqHistory.Chart, chart)
	}

	processor.Storage.SetHistory("nasdaq", nasdaqHistory)

	log.Printf("INFO: nasdaq history loaded. Symbol: %s", content.Symbol)
}
