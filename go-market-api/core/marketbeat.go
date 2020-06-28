package core

import (
	"encoding/json"
	"go-dom-parser/core/parser"
	"go-dom-parser/domain"
	"log"
)

func (processor *Processor) marketBeatParseAnalytics(payload string) {
	var content struct {
		Symbol  string
		Content string
	}

	json.Unmarshal([]byte(payload), &content)

	rows := parser.ParseMarketBeatAnalytics(content.Content)

	rating := domain.Rating{}

	rating.Symbol = content.Symbol

	if len(rows) > 4 {
		if len(rows[1]) > 3 && len(rows[2]) > 3 && len(rows[3]) > 3 && len(rows[4]) > 3 {
			targetPercent, _ := domain.FindPercentValue(rows[5][1])

			rating.ConsensusRating = rows[1][1]
			rating.ConsensusRatingScore = rows[2][1]
			rating.ConsensusBreakdown = rows[3][1]
			rating.TragetPrice = rows[4][1]
			rating.TragetPercent = targetPercent

			// data.Today = []string{"today", rows[1][1], rows[2][1], rows[3][1], rows[4][1], rows[5][1]}
			// data.Days30 = []string{"30 days ago", rows[1][2], rows[2][2], rows[3][2], rows[4][2], rows[5][2]}
			// data.Days90 = []string{"90 days ago", rows[1][3], rows[2][3], rows[3][3], rows[4][3], rows[5][3]}
			// data.Days180 = []string{"180 days ago", rows[1][4], rows[2][4], rows[3][4], rows[4][4], rows[5][4]}

			markerTargetPercent := domain.Marker{
				Symbol: content.Symbol,
				FValue: targetPercent,
				BValue: targetPercent > 50,
			}

			processor.Storage.SetMarker("targetPercent", markerTargetPercent)
		}
	}

	processor.Storage.SetRating("marketbeat", rating)

	log.Printf("INFO: marketbeat rating loaded. Symbol: %s", content.Symbol)
}
