package core

import (
	"encoding/json"
	"fmt"
	"log"
)

type Processor struct {
	ProcessorChan chan string
}

func New() *Processor {
	out := make(chan string)

	return &Processor{
		ProcessorChan: out,
	}
}

// Run - ...
func (p *Processor) Run() {
	go func() {
		for {
			msg := <-p.ProcessorChan
			fmt.Println("INFO: " + msg)

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

			out, err := json.Marshal(result)
			if err != nil {
				log.Printf("ERROR: fail marshl: %s", err.Error)
			}

			// ...

			p.ProcessorChan <- string(out)
		}
	}()
}

type PayloadRMQ struct {
	Content string
	Type    string
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

	rows := parseTable(payload.Content)

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
