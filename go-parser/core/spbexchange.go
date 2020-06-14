package core

import (
	"encoding/json"
	"go-dom-parser/api/sockets"
	"log"
)

func spbParseForeignShare(payload *sockets.InPayload) *sockets.OutPayload {
	// ... parse table

	rows := parseSpbSrc(payload.Content)

	data := make([]SPBStock, len(rows))
	for index, row := range rows {
		if row == nil {
			continue
		}
		data[index].Id = row[0]
		data[index].Marker = row[1]
		data[index].Title = row[2]
		data[index].Code1 = row[3]
		data[index].Code2 = row[4]
		data[index].Count = row[5]
		data[index].Price = row[6]
		data[index].Currency = row[7]
		data[index].Date = row[8]
		data[index].Note = row[9]
	}

	out, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: fail marshl: %s", err.Error)
	}

	result := &sockets.OutPayload{
		Type:    "result.spb.foreignexchange",
		Content: string(out),
	}

	return result
}
