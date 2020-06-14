package core

import (
	"encoding/json"
	"go-dom-parser/api/sockets"
	"log"
)

func marketBeatParseRating(payload *sockets.InPayload) *sockets.OutPayload {
	rows := parseMarketBeatSrc(payload.Content)

	data := MarketBeatStock{}

	// log.Printf("INFO: setup marker %s", payload.Marker)

	data.Marker = payload.Marker
	if len(rows) > 4 {
		if len(rows[1]) > 3 && len(rows[2]) > 3 && len(rows[3]) > 3 && len(rows[4]) > 3 {
			data.Today = []string{"today", rows[1][1], rows[2][1], rows[3][1], rows[4][1], rows[5][1]}
			data.Days30 = []string{"30 days ago", rows[1][2], rows[2][2], rows[3][2], rows[4][2], rows[5][2]}
			data.Days90 = []string{"90 days ago", rows[1][3], rows[2][3], rows[3][3], rows[4][3], rows[5][3]}
			data.Days180 = []string{"180 days ago", rows[1][4], rows[2][4], rows[3][4], rows[4][4], rows[5][4]}
		}
	}

	out, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: fail marshl: %s", err.Error)
	}

	result := &sockets.OutPayload{
		Type:    "result.marketbeat.ratings",
		Content: string(out),
	}

	return result
}
