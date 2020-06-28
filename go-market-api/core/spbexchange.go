package core

import (
	"encoding/json"
	"go-dom-parser/core/parser"
	"go-dom-parser/domain"
	"log"
	"strconv"
)

func (processor *Processor) spbExchangeParseSymbol(payload string) {
	var content = ""
	json.Unmarshal([]byte(payload), &content)

	rows := parser.ParseSpbExchangeSymbol(content)

	// if processor.Storage.Symbols == nil {
	// 	processor.Storage.Symbols = make(map[string]*domain.Symbol)
	// }

	for _, row := range rows {
		if row == nil {
			continue
		}

		id, _ := strconv.ParseUint(row[0], 10, 64)

		symbol := domain.Symbol{
			Symbol:      row[1],
			ID:          id,
			Title:       row[2],
			Price:       row[6],
			Currency:    row[7],
			Description: row[9],
			Ratings:     make(map[string]domain.Rating),
			Infos:       make(map[string]domain.Info),
			Summaries:   make(map[string]domain.Summary),
			Indicators:  make(map[string]domain.Indicator),
			Trades:      make(map[string][]domain.Trade),
			Histories:   make(map[string]domain.History),
			Markers:     make(map[string]domain.Marker),
		}
		processor.Storage.SetSymbol(symbol.Symbol, symbol)

		filter := domain.Filter{
			IsActive: false,
			Symbol:   symbol.Symbol,
		}
		processor.Storage.SetFilterSymbol(symbol.Symbol, filter)

		log.Printf("INFO: spbexchange symbol loaded. Symbol: %s", symbol.Symbol)

		// data[index].Id = row[0]
		// data[index].Marker = row[1]
		// data[index].Title = row[2]
		// data[index].Code1 = row[3]
		// data[index].Code2 = row[4]
		// data[index].Count = row[5]
		// data[index].Price = row[6]
		// data[index].Currency = row[7]
		// data[index].Date = row[8]
		// data[index].Note = row[9]
	}

}
