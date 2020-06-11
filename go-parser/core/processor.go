package core

import (
	"encoding/json"
	"fmt"
	"log"
)

type Processor struct {
	ProcessorChan    chan string
	ProcessorChanOut chan string
}

func New() *Processor {
	out := make(chan string)
	cout := make(chan string)

	return &Processor{
		ProcessorChan:    out,
		ProcessorChanOut: cout,
	}
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			fmt.Println("INFO: income message")

			// ...
			// unmarshal and create docMsg
			payload := &PayloadRMQ{}
			err := json.Unmarshal([]byte(msg), payload)
			if err != nil {
				log.Printf("ERROR: fail unmarshl: %s", err.Error)
			}

			var result = &PayloadRMQ{}
			if payload.Type == "resource.spb" {
				result = parseSPB(payload)
			}

			if payload.Type == "resource.marketbeat" {
				result = parseMarketBeat(payload)
			}

			out, err := json.Marshal(result)
			if err != nil {
				log.Printf("ERROR: fail marshl: %s", err.Error)
			}

			// ...

			fmt.Println("INFO: send processed result to channel")

			p.ProcessorChanOut <- string(out)
		}
	}()
}

type PayloadRMQ struct {
	Content string
	Type    string
	Marker  string
}

type SPBStock struct {
	Id       string
	Marker   string
	Title    string
	Code1    string
	Code2    string
	Count    string
	Price    string
	Currency string
	Date     string
	Note     string
}

func parseSPB(payload *PayloadRMQ) *PayloadRMQ {
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

	//fmt.Println("####### rows = ", len(rows), rows)
	// ... parse table

	//log.Printf("ERROR: fail unmarshl: %s", doc)

	result := &PayloadRMQ{
		Type:    "result.spb",
		Content: string(out),
	}

	return result
}

type MarketBeatStock struct {
	Marker  string
	Today   []string
	Days30  []string
	Days90  []string
	Days180 []string
}

func parseMarketBeat(payload *PayloadRMQ) *PayloadRMQ {
	// ... parse table

	rows := parseMarketBeatSrc(payload.Content)

	// data := make([]MarketBeatStock, 5)

	data := MarketBeatStock{}

	log.Printf("INFO: setup marker %s", payload.Marker)

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

	//fmt.Println("####### rows = ", len(rows), rows)
	// ... parse table

	//log.Printf("ERROR: fail unmarshl: %s", doc)

	result := &PayloadRMQ{
		Type:    "result.marketbeat",
		Content: string(out),
	}

	return result
}
